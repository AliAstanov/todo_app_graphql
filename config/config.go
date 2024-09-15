package config

type PsqlConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type Config struct {
	PsqlConfig PsqlConfig
}

func Load() Config {

	var cfg Config

	cfg.PsqlConfig.User = "boot"
	cfg.PsqlConfig.Password = "root"
	cfg.PsqlConfig.Host = "localhost"
	cfg.PsqlConfig.Port = 5432
	cfg.PsqlConfig.Database = "exam"

	return cfg
}
