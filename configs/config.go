package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RedisAddr      string `mapstructure:"REDIS_ADDR"`
	RateLimitIP    int    `mapstructure:"RATE_LIMIT_IP"`
	RateLimitToken int    `mapstructure:"RATE_LIMIT_TOKEN"`
	BlockTime      int    `mapstructure:"BLOCK_TIME"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
