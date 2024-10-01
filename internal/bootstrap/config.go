package bootstrap

import (
	"time"

	"github.com/spf13/viper"
)

type Cookie struct {
	Name     string        `mapstructure:"name"`
	Path     string        `mapstructure:"path"`
	MaxAge   time.Duration `mapstructure:"max_age"`
	HttpOnly bool          `mapstructure:"http_only"`
	Secure   bool          `mapstructure:"secure"`
}

type Session struct {
	Cookie Cookie `mapstructure:"cookie"`
}

type Server struct {
	Port             int      `mapstructure:"port"`
	Session          Session  `mapstructure:"session"`
	CorsAllowOrigins []string `mapstructure:"cors_allow_origins"`
}

type Database struct {
	URL string `mapstructure:"url"`
}

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

func Setup(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
