package main

import (
	"github.com/Herts/ray-ui/communication"
	"github.com/Herts/ray-ui/models"
	_ "github.com/Herts/ray-ui/routers"
	"github.com/astaxie/beego"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	models.InitDb()
}

func timelyTask() {
	interval := beego.AppConfig.DefaultInt64("dataupdateinterval", 60)
	gocron.
		Every(uint64(interval)).
		Minutes().
		From(gocron.NextTick()).
		Do(communication.UpdateAllUserDataOnAllServers,
			time.Now().Add(-24*time.Hour).Format("2006-01-02"),
			"")
	userInterval := beego.AppConfig.DefaultInt64("userupdateinterval", 24)
	gocron.
		Every(uint64(userInterval)).
		Hours().
		From(gocron.NextTick()).
		Do(communication.RetrieveAllUserOnAllServers)

	gocron.
		Every(uint64(interval)).
		Minutes().
		From(gocron.NextTick()).
		Do(models.SumUserData)

	gocron.Start()
}

func main() {
	go timelyTask()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
