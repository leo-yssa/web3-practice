package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Version  string          `yaml:"version"`
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"database"`
	Cache    *CacheConfig    `yaml:"cache"`
	Jwt      *JwtConfig      `yaml:"jwt"`
	Oauth    *OauthConfig    `yaml:"oauth"`
	Gateway  *GatewayConfig  `yaml:"gateway"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Kind   string `yaml:"kind"`
	User   string `yaml:"user"`
	Secret string `yaml:"secret"`
	Host   string `yaml:"host"`
	Schema string `yaml:"schema"`
}

type CacheConfig struct {
	Kind   string `yaml:"kind"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

type JwtConfig struct {
	Secret  string          `yaml:"secret"`
	Access  *JwtTokenConfig `yaml:"access"`
	Refresh *JwtTokenConfig `yaml:"refresh"`
}

type JwtTokenConfig struct {
	Duration string `yaml:"duration"`
}

type OauthConfig struct {
	Google GoogleConfig `yaml:"google"`
}

type GoogleConfig struct {
	Client   GoogleClient   `yaml:"client"`
	Redirect GoogleRedirect `yaml:"redirect"`
}

type GoogleClient struct {
	Id     string `yaml:"id"`
	Secret string `yaml:"secret"`
}

type GoogleRedirect struct {
	Login   string `yaml:"login"`
	Youtube string `yaml:"youtube"`
}

type GatewayConfig struct {
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
