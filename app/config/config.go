package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"DB_SSLMODE"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}