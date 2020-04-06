package main

import (
	"./models"
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	models.InitDb()
}

func main() {
	beego.Run()
}
