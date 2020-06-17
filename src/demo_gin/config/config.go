package config

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	//app
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
	//email
	MailUser         string `json:"mail_user"`
	MailUserPassword string `json:"mail_password"`
	MailUserHost     string `json:"mail_host"`
	MailUserPort     string `json:"mail_port"`
	//db
	DbDriver string `json:"db_driver"`
	DbHost   string `json:"db_host"`
	DbPort   string `json:"db_port"`
	DbUser   string `json:"db_user"`
	DbPass   string `json:"db_pass"`
	DbDb     string `json:"db_db"`
	//redis
	RedisHost string `json:"redis_host"`
	RedisPort string `json:"redis_port"`
	RedisUser string `json:"redis_user"`
	RedisPass string `json:"redis_pass"`
}

type MailConfig struct {
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func init() {
	file, err := os.Open("./config/app.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_cfg); err != nil {
		panic(err)
	}
}
