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

func AddUserServer(us *UserServer) {
	db.Save(us)
}
