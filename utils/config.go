package utils

import "github.com/spf13/viper"

type Config struct {
	Token string `mapstructure:"BOT_TOKEN"`
}

func LoadEnv() Config {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)

	return config
}
