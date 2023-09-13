package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App   string `yaml:"app"`
	Port  int    `yaml:"port"`
	MySQL *MySQL `yaml:"mysql"`
	Redis *Redis `yaml:"redis"`
	K8s   *K8s   `yaml:"k8s"`
}

type MySQL struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	prefix       string `yaml:"prefix"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	LogMode      bool   `yaml:"log_mode"`
	LogZap       bool   `yaml:"log_zap"`
}

type Redis struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	Password         string `yaml:"password"`
	DB               int    `yaml:"db"`
	PoolSize         int    `yaml:"pool_size"`
	IdleTimeout      int    `yaml:"idle_timeout"`
	MaxIdleTimeout   int    `yaml:"max_idle_timeout"`
	MaxActiveTimeout int    `yaml:"max_active_timeout"`
}

type K8s struct {
	Config string `yaml:"config"`
}

var globalConfig *Config

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	err = viper.Unmarshal(&globalConfig)
	return globalConfig, err
}

func GetConfig() *Config {
	return globalConfig
}
