package config

import "github.com/spf13/viper"

type Config struct {
	Dbname   string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Password string `mapstructure:"DB_PASSWORD"`
	AppPort  int    `mapstructure:"APP_PORT"`
}

func ReadConfig(path string, conf *Config) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("bankapp")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
