package config

type RedisConfig struct {
	Addr     string `mapstructure:"address"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}
