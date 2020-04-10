package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RemoteServer struct {
	gorm.Model    `json:"-"`
	NickName      string    `json:"nickName"`
	Address4      string    `json:"address4"`
	ServerName    string    `json:"serverName"`
	Host          string    `json:"host"`
	TLSName       string    `json:"tlsName"`
	Port          int       `json:"port"`
	Provider      string    `json:"provider"`
	Price         float64   `json:"price"`
	Region        string    `gorm:"primary_key" json:"region"`
	Index         int       `gorm:"primary_key;auto_increment:false" json:"index"`
	StreamSetting string    `json:"streamSetting"`
	Path          string    `json:"path"`
	ConfsType     string    `json:"confsType"`
	Tls           string    `json:"tls"`
	ExpiresOn     time.Time `json:"expiresOn"`
	ApiKey        string    `json:"apiKey"`
}

func AddServer(server *RemoteServer) {
	db.Save(server)
}

func GetAllServers() (servers []*RemoteServer) {
	db.Find(&servers)
	return
}

func GetServer(server *RemoteServer) *RemoteServer {
	var s RemoteServer
	db.Where(server).First(&s)
	return &s
}
