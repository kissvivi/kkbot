package kkbot

import (
	"github.com/spf13/viper"
)

type Config struct {
	App        App        `json:"app"`
	DataSource DataSource `json:"datasource"`
}

type App struct {
}

type DataSource struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type Mysql struct {
	//Host     string `json:"host"`
	//Port     string `json:"port"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type Redis struct {
	//Host     string `json:"host"`
	//Port     string `json:"port"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func InitConfig() (*Config, error) {
	v := viper.New()

	//v.AddConfigPath(".")
	//v.SetConfigName("config")
	v.SetConfigFile("config.yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
