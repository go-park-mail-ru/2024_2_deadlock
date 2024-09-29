package bootstrap

import "github.com/spf13/viper"

type Database struct {
	URL string `mapstructure:"url"`
}

type Config struct {
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
