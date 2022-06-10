package config

import (
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
	Develop  bool   `mapstructure:"develop"`
	LogLevel string `mapstructure:"log_level"`
	Server   Server `mapstructure:"server"`
	PG       PG     `mapstructure:"pg"`
}

// ReadConfig reads base config and merges config if overwritePath is givven.
func ReadConfig(
	basePath string,
	overwritePath string,
) (*Config, error) {
	viper.SetConfigFile(basePath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error reading base config: "+basePath)
	}

	if overwritePath != "" {
		viper.SetConfigFile(overwritePath)

		err := viper.MergeInConfig()
		if err != nil {
			return nil, errors.Wrap(err, "error merging with config: "+overwritePath)
		}
	}

	conf := new(Config)

	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling config")
	}

	return conf, nil
}
