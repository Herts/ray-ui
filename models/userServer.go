package models

import "github.com/jinzhu/gorm"

type UserServer struct {
	gorm.Model       `json:"-"`
	Email            string  `gorm:"primary_key" json:"email"`
	NickName         string  `json:"nickName"`
	Region           string  `gorm:"primary_key" json:"region"`
	Index            int     `gorm:"primary_key;auto_increment:false" json:"index"`
	UpDataConsumed   float64 `json:"upDataConsumed"`
	DownDataConsumed float64 `json:"downDataConsumed"`
	UserId           string  `json:"userId"`
	Enabled          bool    `json:"enabled"`
	Level            int     `json:"level,string"`
	AlterID          int     `json:"alterId,string"`
}

func GetAllUserServers() (userServers []*UserServer) {
	db.Find(&userServers)
	return
}

func SaveUserServer(us *UserServer) {
	db.Save(us)
}

func GetUserOnServer(email string, region string, index int) *UserServer {
	var userServer UserServer
	db.FirstOrInit(&userServer, UserServer{Email: email,
		Region: region,
		Index:  index})
	return &userServer
}
