package config

import (
    "github.com/Unknwon/goconfig"
)

var Config config

type config struct {
    JwtAuthKey     string
    Port           string
    EmailKey       string
    DBUser         string
    DBPassword     string
    DBName         string
    DBAddr         string
    Admin          string
    PWSecret       string
}

func init() {
    configFile, err := goconfig.LoadConfigFile("config/config.ini")
    if err != nil {
        panic(err)
    }
    Config.JwtAuthKey, err = configFile.GetValue("Secret", "JwtAuthKey")
    if err != nil {
        panic(err)
    }
}