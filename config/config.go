package config

import (
	"flag"

	"github.com/spf13/viper"
)

var ServerMode = flag.String("mode", "dev", "运行环境")

type Config struct {
	Server ServerConfig   `mapstructure:"Server"`
	DB     DataBaseConfig `mapstructure:"DataBase"`
	Auth   AuthConfig     `mapstructure:"Auth"`
	Logger LoggerConfig   `mapstructure:"Logger"`
	Redis  RedisConfig    `mapstructure:"Redis"`
}

func NewConfig(path string) (config Config, err error) {
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	viper.AddConfigPath(".")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath(path)

	viper.SetConfigName("config." + *ServerMode)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
