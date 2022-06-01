package config

type ServerConfig struct {
	AppName    string `mapstructure:"app_name"`
	Port       string `mapstructure:"port"`
	DBDriver   string `mapstructure:"db_driver" yaml:"db_driver"`
	UploadPath string `mapstructure:"upload_path"`
}
