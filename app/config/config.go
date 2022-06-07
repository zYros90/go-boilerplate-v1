package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Auth struct {
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type PG struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DBName string `mapstructure:"db_name"`
	SSL    string `mapstructure:"ssl"`
	Auth   Auth   `mapstructure:"auth"`
}

type Server struct {
	Host         string   `mapstructure:"host"`
	Port         int      `mapstructure:"port"`
	JWTSecret    string   `mapstructure:"jwtSecret"`
	AllowOrigins []string `mapstructure:"allow_origins"`
}

type Config struct {
	Develop bool   `mapstructure:"develop"`
	Debug   bool   `mapstructure:"debug"`
	Server  Server `mapstructure:"server"`
	PG      PG     `mapstructure:"pg"`
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfig(
	basePath string,
	overwritePath string,
) (*Config, error) {
	conf := new(Config)
	viper.SetConfigFile(basePath)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.Wrap(err, "couldn't find base config: "+basePath)
		} else {
			return nil, errors.Wrap(err, "error reading base config: "+basePath)
		}
	}

	if overwritePath != "" {
		fmt.Println("overwriting config with ", overwritePath)
		viper.SetConfigFile(overwritePath)
		err := viper.MergeInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				return nil, errors.Wrap(err, "couldn't find merge config: "+basePath)
			} else {
				return nil, errors.Wrap(err, "error reading merge config: "+basePath)
			}
		}
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}
	return conf, nil
}
