package config

import (
	"github.com/spf13/viper"
	"time"
)

// Database holds the database configuration
type Database struct {
	Host                  string
	Port                  int
	Username              string
	Password              string
	Name                  string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
}

var db Database

// DB returns the default database configuration
func DB() Database {
	return db
}

// LoadDB loads database configuration
func LoadDB() {
	mu.Lock()
	defer mu.Unlock()

	db = Database{
		Name:                  viper.GetString("database.name"),
		Username:              viper.GetString("database.username"),
		Password:              viper.GetString("database.password"),
		Host:                  viper.GetString("database.host"),
		Port:                  viper.GetInt("database.port"),
		MaxOpenConnections:    viper.GetInt("database.max_open_connections"),
		MaxIdleConnections:    viper.GetInt("database.max_idle_connections"),
		MaxConnectionLifetime: viper.GetDuration("database.max_connection_lifetime") * time.Minute,
	}
}
