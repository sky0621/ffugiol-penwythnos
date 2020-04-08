package sys

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env        string `default:"local"`
	DBHost     string `split_words:"true" default:"localhost"`
	DBPort     string `split_words:"true" default:"65432"`
	DBName     string `split_words:"true" default:"localdb"`
	DBUser     string `split_words:"true" default:"postgres"`
	DBPassword string `split_words:"true" default:"localpass"`
	DBSSLMode  string `envconfig:"db_sslmode" split_words:"true" default:"disable"`
	ServerPort string `split_words:"true" default:":8080"`
}

func InitConfig() *Config {
	var cfg Config
	envconfig.MustProcess("FS", &cfg)
	return &cfg
}
