package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want *env
	}{
		{
			name: "def",
			env:  nil,
			want: &env{
				Env:                     "local",
				ServerPort:              "5050",
				DBDriverName:            "postgres",
				DBHost:                  "localhost",
				DBPort:                  "19999",
				DBName:                  "localdb",
				DBUser:                  "postgres",
				DBPassword:              "localpass",
				DBSSLMode:               "disable",
				MovieBucket:             "bucket",
				AuthDebug:               true,
				AuthCredentialsOptional: true,
				Auth0Domain:             "localhost",
				Auth0ClientID:           "authid",
				Auth0ClientSecret:       "authsecret",
				Auth0Audience:           "http://api.localhost",
			},
		},
		{
			name: "set",
			env: map[string]string{
				"FS_ENV":                       "dev",
				"FS_SERVER_PORT":               "8080",
				"FS_DB_DRIVER_NAME":            "mysql",
				"FS_DB_HOST":                   "example.com",
				"FS_DB_PORT":                   "3456",
				"FS_DB_NAME":                   "db01",
				"FS_DB_USER":                   "testuser",
				"FS_DB_PASSWORD":               "passpass",
				"FS_DB_SSL_MODE":               "true",
				"FS_MOVIE_BUCKET":              "bucket",
				"FS_AUTH_DEBUG":                "false",
				"FS_AUTH_CREDENTIALS_OPTIONAL": "false",
				"FS_AUTH0_DOMAIN":              "example2.com",
				"FS_AUTH0_CLIENT_ID":           "authidid",
				"FS_AUTH0_CLIENT_SECRET":       "authsecsec",
				"FS_AUTH0_AUDIENCE":            "https://example.com",
			},
			want: &env{
				Env:                     "dev",
				ServerPort:              "8080",
				DBDriverName:            "mysql",
				DBHost:                  "example.com",
				DBPort:                  "3456",
				DBName:                  "db01",
				DBUser:                  "testuser",
				DBPassword:              "passpass",
				DBSSLMode:               "true",
				MovieBucket:             "bucket",
				AuthDebug:               false,
				AuthCredentialsOptional: false,
				Auth0Domain:             "example2.com",
				Auth0ClientID:           "authidid",
				Auth0ClientSecret:       "authsecsec",
				Auth0Audience:           "https://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				if err := os.Setenv(k, v); err != nil {
					t.Fail()
				}
			}
			if got := loadEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
