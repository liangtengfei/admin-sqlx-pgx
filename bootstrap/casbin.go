package bootstrap

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/redis-adapter/v2"
	"log"
	"study.com/demo-sqlx-pgx/config"
)

func LoadCasbin(cfg config.Config) *casbin.Enforcer {
	a := redisadapter.NewAdapter("tcp", cfg.Redis.Addr)
	e, err := casbin.NewEnforcer("resources/rbac_model.conf", a)
	if err != nil {
		log.Fatal("获取权限控制文件异常：", err)
	}
	e.EnableLog(true)
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal("加载权限控制策略异常：", err)
	}
	if ok := e.HasPolicy("alice", "data1", "POST"); !ok {
		_, err := e.AddPolicy("alice", "data1", "POST")
		if err != nil {
			log.Fatal("新增权限控制策略异常：", err)
		}
	}

	// Save the policy back to DB.
	//err = e.SavePolicy()
	//if err != nil {
	//	log.Fatal("保存权限控制策略异常：", err)
	//}
	return e
}
