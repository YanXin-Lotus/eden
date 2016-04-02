package config

import (
    "github.com/astaxie/beego"
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
}

func init() {
    configFile, err := goconfig.LoadConfigFile("config/config.ini")
    if err != nil {
        beego.Error(err)
        panic(err)
    }
    Config.JwtAuthKey, err = configFile.GetValue("Key", "JwtAuthKey")
    if err != nil {
        beego.Error(err)
        panic(err)
    }
}