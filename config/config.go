// Package config provides a centralized configuration management system for the application.
// It utilizes the Viper library to load and manage configuration settings from a YAML file
// and environment variables. The configuration is loaded only once in a thread-safe manner
// using the sync.Once mechanism, ensuring efficient access throughout the application.
//
// The configuration is structured into two main sections:
//   - Db: Contains the database connection settings, including host, port, user credentials,
//     database name, SSL mode, and time zone.
//   - Server: Contains the server-related settings, such as the port number on which the server runs.
//
// The GetConfig function initializes and returns the application configuration. It reads the
// configuration from a YAML file named "config.yaml" located in the application's root directory.
// Additionally, environment variables can be used to override the YAML configuration values.
// Environment variables are expected to follow the format CONFIG_KEY where key is the YAML key
// with periods replaced by underscores.
//
// Example usage:
//
//	config := config.GetConfig()
//	fmt.Println("Server running on port:", config.Server.Port)
//	fmt.Println("Database host:", config.Db.Host)
//
// The package should be imported and used wherever configuration settings are needed in the application.
package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type Config struct {
	Db     *db
	Server *server
}

type server struct {
	Port int
}

type db struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
