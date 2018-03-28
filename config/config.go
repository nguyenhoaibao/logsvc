package config

import (
	"io"
	"strings"

	"github.com/spf13/viper"
)

const (
	envPrefix = "LOGSVC"
	confType  = "YML"
)

// Config is a struct represents the service config.
type Config struct {
	Server struct {
		Addr string
	}
	Database struct {
		Addr     string
		User     string
		Password string
		Name     string
	}
}

// Load loads config from io.Reader and/or environment variables.
func Load(r io.Reader) (*Config, error) {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	_ = viper.BindEnv("server.addr")
	_ = viper.BindEnv("database.addr")
	_ = viper.BindEnv("database.user")
	_ = viper.BindEnv("database.password")
	_ = viper.BindEnv("database.name")

	viper.SetConfigType(confType)
	if err := viper.MergeConfig(r); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
