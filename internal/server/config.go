package server

import (
	"github.com/spf13/viper"
)

type Config struct {
	Version  string         `yaml:"version"`
	Server   serverConfig   `yaml:"server"`
	Database databaseConfig `yaml:"database"`
	Cache    cacheConfig    `yaml:"cache"`
	Gateway  gatewayConfig  `yaml:"gateway"`
}

type serverConfig struct {
	Port string `yaml:"port"`
}

type databaseConfig struct {
	Kind   string `yaml:"kind"`
	User   string `yaml:"user"`
	Secret string `yaml:"secret"`
	Host   string `yaml:"host"`
	Schema string `yaml:"schema"`
}

type cacheConfig struct {
	Kind   string `yaml:"kind"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

type gatewayConfig struct {
	Host string `yaml:"host"`
}

func InitConfig(name string) (*Config, error) {
	v := initViper(name)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var c *Config
	err = viper.Unmarshal(&c)
	return c, err
}

func initViper(name string) *viper.Viper {
	v := viper.GetViper()
	v.AddConfigPath(".")
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()
	v.SetConfigName(name)
	return v
}
