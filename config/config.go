package config

type Config struct {
	parseUrl string
}

func NewConfig() *Config {
	return &Config{
		parseUrl: "",
	}
}