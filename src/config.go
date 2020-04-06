package system

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env        string `default:"local"`
	DBHost     string `split_words:"true" default:"localhost"`
	DBPort     string `split_words:"true" default:"65432"`
	DBName     string `split_words:"true" default:"localdb"`
	DBUser     string `split_words:"true" default:"postgres"`
	DBPassword string `split_words:"true" default:"localpass"`
	ListenPort string `split_words:"true" default:":5432"`
}

func LoadConfig() Config {
	var cfg Config
	envconfig.Process("fs", &cfg)
	return cfg
}
