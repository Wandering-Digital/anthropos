package config

import (
	"time"

	"github.com/spf13/viper"
)

// Database holds the database configuration
type Database struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Name            string
	Options         map[string][]string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	BatchSize       int
}

var (
	db Database
)

// DB returns the default database configuration
func DB() *Database {
	return &db
}

// LoadDB loads database configuration
func loadDB() {
	db = Database{
		Name:            viper.GetString("database.name"),
		Username:        viper.GetString("database.username"),
		Password:        viper.GetString("database.password"),
		Host:            viper.GetString("database.host"),
		Port:            viper.GetInt("database.port"),
		Options:         viper.GetStringMapStringSlice("database.options"),
		MaxIdleConn:     viper.GetInt("database.max_idle_connection"),
		MaxOpenConn:     viper.GetInt("database.max_open_connection"),
		MaxConnLifetime: viper.GetDuration("database.max_connection_lifetime") * time.Second,
		BatchSize:       viper.GetInt("database.batch_size"),
	}
}
