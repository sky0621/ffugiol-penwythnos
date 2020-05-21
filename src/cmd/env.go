package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type env struct {
	Env        string `default:"local"`
	ServerPort string `split_words:"true" default:"5050"`
	// DB接続関連
	DBDriverName string `split_words:"true" default:"postgres"`
	DBHost       string `split_words:"true" default:"localhost"`
	DBPort       string `split_words:"true" default:"19999"`
	DBName       string `split_words:"true" default:"backenddb"`
	DBUser       string `split_words:"true" default:"postgres"`
	DBPassword   string `split_words:"true" default:"localpass"`
	DBSSLMode    string `envconfig:"db_ssl_mode" split_words:"true" default:"disable"`
	// GCS関連
	MovieBucket string `split_words:"true" default:"bucket"`
	// PubSub関連
	CreateMovieTopic string `split_words:"true" default:"create_movie_topic"`
	// Auth関連
	AuthDebug               bool   `split_words:"true" default:"true"`
	AuthCredentialsOptional bool   `split_words:"true" default:"true"`
	Auth0Domain             string `split_words:"true" default:"localhost"`
	Auth0ClientID           string `split_words:"true" default:"authid"`
	Auth0ClientSecret       string `split_words:"true" default:"authsecret"`
	Auth0Audience           string `split_words:"true" default:"http://api.localhost"`
	Auth0Timeout            string `split_words:"true" default:"30s"`
}

func loadEnv() *env {
	var e env
	if err := envconfig.Process("fs", &e); err != nil {
		fmt.Println(err)
		return nil
	}
	return &e
}
