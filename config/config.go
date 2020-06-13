package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config add your desired environment variable here
type Config struct {
	AppName         string `envconfig:"app_name"`
	AppQuote        string `envconfig:"app_quote"`
	LogLevel        string `envconfig:"log_level"`
	MongoDbName     string `envconfig:"mongo_db_name"`
	MongoDbDsn      string `envconfig:"mongo_db_dns"`
	SandboxAccount  string `envconfig:"sandbox_account"`
	SandboxEmail    string `envconfig:"sandbox_email"`
	SandboxPassword string `envconfig:"sandbox_password"`
	SandboxRole     string `envconfig:"sandbox_role"`
	ScheduleEvery   string `envconfig:"schedule_every"`
	ScheduleTime    string `envconfig:"schedule_time"`
	Version         string `envconfig:"version"`
}

// C ...
var C Config

func init() {
	err := envconfig.Process("angelica_worker", &C)
	if err != nil {
		log.Fatal(err.Error())
	}
}
