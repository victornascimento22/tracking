package configs

import "github.com/spf13/viper"

var cfg *conf

type conf struct {
	Url string `mapstructure:"URL"`
}

func LoadConfig(path string) (*conf, error) {

	viper.SetConfigName("app_config")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {

		panic(err)

	}

	err = viper.Unmarshal(&cfg)
	if err != nil {

		panic(err)

	}

	return cfg, nil
}
