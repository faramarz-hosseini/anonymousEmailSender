package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Email string `mapstructure:"email"`
}

var defaults = map[string]string{
	"email": "test@test.com",
}

func LoadConfig(configFile string) (*Config, error) {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("ANONYMOUS_EMAIL_SENDER")
	viper.AutomaticEnv()

	var config Config
	if configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
