package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RedisAddr    string   `yaml:"redis_addr"`
	MySQLDSN     string   `yaml:"mysql_dsn"`
	KafkaBrokers []string `yaml:"kafka_brokers"`
}

func Load() *Config {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
