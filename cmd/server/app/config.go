package app

type Database struct {
	URL string `mapstructure:"url"`
}

type Config struct {
	Database Database `mapstructure:"database"`
}
