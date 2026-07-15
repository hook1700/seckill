package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Redis RedisConfig `mapstructure:"redis"`
	MySQL MySQLConfig `mapstructure:"mysql"`
	Kafka KafkaConfig `mapstructure:"kafka"`
}

type AppConfig struct {
	Port       int `mapstructure:"port"`
	GOMAXPROCS int `mapstructure:"gomaxprocs"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	PoolSize int    `mapstructure:"pool_size"`
}

type MySQLConfig struct {
	DSN          string `mapstructure:"dsn"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type KafkaConfig struct {
	Brokers       string `mapstructure:"brokers"`
	Topic         string `mapstructure:"topic"`
	ConsumerGroup string `mapstructure:"consumer_group"`
}

func Load() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")        // 容器根目录
	v.AddConfigPath("./config") // 可选

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}
	return &cfg
}
