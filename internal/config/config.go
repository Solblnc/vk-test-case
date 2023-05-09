package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Responses Responses
}

type Responses struct {
	Start     string `mapstructure:"start"`
	Choice    string `mapstructure:"choice"`
	MadLads   string `mapstructure:"mad_lads"`
	SolCasino string `mapstructure:"solcasino"`
	Lily      string `mapstructure:"lily"`
	OkayBears string `mapstructure:"okay_bears"`
}

// Init - reading yml file and returns struct
func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.UnmarshalKey("messages", &cfg.Responses); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// FromEnv - reads an environment variables
func FromEnv(token string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(token).(string)
	if !ok {
		log.Fatal(err)
	}
	return value
}
