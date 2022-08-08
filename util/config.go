package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read bt viper from a config file or environment varibales
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBsource             string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
