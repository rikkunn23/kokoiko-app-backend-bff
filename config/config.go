package config

import (
	"os"

	"github.com/caarlos0/env"
)



var conf *Config
type Environment string

const (
	Local Environment = "local"
	Test Environment = "test"
	Development Environment = "dev"
	Stage1 Environment = "stage1"
	Production Environment = "prod"
)

func Get() *Config {
	return conf
}

func GetGoEnv() Environment {
	return Environment(GetGoEnvString())
}

// envに定義したGO_ENVを取得するラッパー関数
func GetGoEnvString() string {
	return os.Getenv("GO_ENV")
}

// GetGoEnvDir ..
func GetGoEnvDir() string {
	return os.Getenv("GO_ENV_DIR")
}

// GetAPIKey ...
func GetAPIKey() string {
	return os.Getenv("API_KEY")
}

// // IsAWSEnv ...
// func IsAWSEnv() bool {
// 	e := GetGoEnv()
// 	return !(e == Local || e == Test)
// }


func NewConfig()error{
	conf = &Config{
		App: &AppConfig{},
		Postgres: &PostgresConfig{},
		HTTPClient: &HTTPClientConfig{},
}
return env.Parse(conf)
}
