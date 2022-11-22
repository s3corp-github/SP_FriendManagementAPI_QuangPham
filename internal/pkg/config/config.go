package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config store all configuration of the application
// The value are read by viper from the congfig file
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

const (
	// ConfigName env file name
	ConfigName = "app"
	// ConfigExtension  env file extension
	ConfigExtension = "env"
)

// LoadConfig config value from env file
func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigExtension)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	if config.DBSource == "" {
		log.Fatal("Cannot load DBSource from config file.")
	}

	return config, nil
}
