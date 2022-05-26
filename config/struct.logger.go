package config

type LoggerConfig struct {
	Level   string `mapstructure:"level"`
	Path    string `mapstructure:"path"`
	MaxSize int    `mapstructure:"max_size"`
	MaxAge  int    `mapstructure:"max_age"`
}
