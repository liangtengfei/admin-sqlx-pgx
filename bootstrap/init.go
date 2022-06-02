package bootstrap

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os/signal"
	"study.com/demo-sqlx-pgx/config"
	"study.com/demo-sqlx-pgx/global"
	"study.com/demo-sqlx-pgx/pkg/cache"
	"study.com/demo-sqlx-pgx/pkg/token"
	"study.com/demo-sqlx-pgx/pkg/zaplog"
	"study.com/demo-sqlx-pgx/router"
	"study.com/demo-sqlx-pgx/service"
	"syscall"
	"time"
)

func attachConfig() config.Config {
	cfg, err := config.NewConfig("./resources")
	if err != nil {
		log.Fatal("配置信息加载错误", err)
	}
	return cfg
}

func attachTokenMaker(symmetricKey string) token.Maker {
	tokenMaker, err := token.NewJWTMaker(symmetricKey)
	if err != nil {
		log.Fatal("加载JWT授权失败", err)
	}
	return tokenMaker
}

func attachLogger(config config.LoggerConfig, mode string) *zap.Logger {
	core := zapcore.NewCore(
		zaplog.GetEncoder(),
		zaplog.GetLumberWriter(config.Path, config.MaxSize, config.MaxAge, mode),
		zaplog.GetLevel(config.Level),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}

func attachDbConn(cfg config.Config) *sqlx.DB {
	conn, err := sql.Open(cfg.Server.DBDriver, cfg.DB.ConnDsn(cfg.Server.DBDriver))
	if err != nil {
		log.Fatal("数据库信息加载错误: ", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal("数据库连接错误，请确认数据库服务是否开启: ", err)
	}
	// 最大打开的连接数，<= 0 不限制
	conn.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	// 连接池大小，默认大小为 2，<= 0 时不使用连接池
	conn.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	// 一个连接可以被重用的最大时限，也就是它在连接池中的最大存活时间，0 表示可以一直重用
	conn.SetConnMaxLifetime(1 * time.Minute)
	conn.SetConnMaxIdleTime(2 * time.Minute)

	return sqlx.NewDb(conn, cfg.Server.DBDriver)
}

func RunServer(staticFs embed.FS) {
	//运行模式 默认：dev
	mode := "DEBUG"
	if *config.ServerMode == "prod" {
		mode = "PROD"
	}

	// 加载配置信息
	cfg := attachConfig()
	global.Config = cfg

	// 加载Token生成
	tokenMaker := attachTokenMaker(cfg.Auth.TokenSymmetricKey)
	global.TokenMaker = tokenMaker

	// 加载权限控制策略
	enforcer := LoadCasbin(cfg)
	global.Enforcer = enforcer

	// 配置日志
	logger := attachLogger(cfg.Logger, mode)
	defer logger.Sync()
	global.Log = logger

	// 加载缓存配置
	cacheStore := cache.NewRedisCache(cfg)
	global.CacheStore = cacheStore

	// 数据库连接
	sqlxDB := attachDbConn(cfg)
	// 注入数据库
	service.InitService(sqlxDB)

	// 缓存参数配置信息
	err := service.ConfigCacheAll()
	if err != nil {
		log.Fatal("缓存参数配置错误：", err)
	}

	// 加载gin引擎
	engine := router.InitRouter(staticFs)
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: engine,
	}

	log.Println(fmt.Sprintf("服务启动成功：http://%s", "127.0.0.1"+cfg.Server.Port))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅关闭
	gracefullyShutdown(srv)
}

func gracefullyShutdown(srv *http.Server) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("关闭苏服务错误", err)
	}
}
