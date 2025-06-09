package config

import (
	"time"

	"github.com/spf13/viper"
)

type LogLevel string
type Environment string

const (
	EnvironmentLocal       Environment = "local"
	EnvironmentDevelopment Environment = "development"
	EnvironmentStaging     Environment = "staging"
	EnvironmentProduction  Environment = "production"
)

type Application struct {
	Base         string
	Env          Environment
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var app Application

func App() *Application {
	return &app
}

func loadApp() {
	env := EnvironmentDevelopment
	if e := Environment(viper.GetString("app.env")); e != "" {
		env = e
	}

	app = Application{
		Base:         viper.GetString("app.host"),
		Port:         viper.GetInt("app.port"),
		Env:          env,
		ReadTimeout:  viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("app.write_timeout") * time.Second,
		IdleTimeout:  viper.GetDuration("app.idle_timeout") * time.Second,
	}
}
