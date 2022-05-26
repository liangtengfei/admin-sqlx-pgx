package config

import "fmt"

type DataBaseConfig struct {
	Path         string `mapstructure:"path"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"db_name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

func (d *DataBaseConfig) ConnDsn(driver string) string {
	if driver == "mysql" {
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&tls=false&loc=Local",
			d.Username,
			d.Password,
			d.Path,
			d.Port,
			d.DbName,
		)
	} else if driver == "postgresql" {
		return fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			d.Username,
			d.Password,
			d.Path,
			d.Port,
			d.DbName,
		)
	} else if driver == "pgx" {
		return fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			d.Username,
			d.Password,
			d.Path,
			d.Port,
			d.DbName,
		)
	}
	return ""
}

func (d *DataBaseConfig) ConnDsnMySQL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&tls=false&loc=Local",
		d.Username,
		d.Password,
		d.Path,
		d.Port,
		d.DbName,
	)
}

func (d *DataBaseConfig) ConnDsnPostgresQL() string {
	// postgresql://postgres:xiaohuozhi@localhost:5432/tengfeiliang?sslmode=disable
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		d.Username,
		d.Password,
		d.Path,
		d.Port,
		d.DbName,
	)
}
