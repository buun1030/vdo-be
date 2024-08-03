package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Load(getEnv func(string) string) (*Config, error) {
	env := strings.ToLower(getEnv("ENV"))
	if env == "" {
		env = "local"
	}

	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	viper.SetConfigName(env)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	AppConfig *AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Env       string `mapstructure:"env"`
	AppDomain string `mapstructure:"app_domain"`
	ApiDomain string `mapstructure:"api_domain"`
}
