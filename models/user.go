package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model       `json:"-"`
	Email            string `gorm:"primary_key" json:"email"`
	UserId           string `json:"userId"`
	UpDataConsumed   int64  `json:"upDataConsumed"`
	DownDataConsumed int64  `json:"downDataConsumed"`
	AccessToken      string `json:"accessToken"`
}

type UserData struct {
	gorm.Model       `json:"-"`
	Region           string `gorm:"primary_key" json:"region"`
	Index            int    `gorm:"primary_key;auto_increment:false" json:"index"`
	Email            string `gorm:"primary_key" json:"email"`
	Date             string `gorm:"primary_key;date" json:"date"`
	UpDataConsumed   int64  `json:"upDataConsumed"`
	DownDataConsumed int64  `json:"downDataConsumed"`
}

func AddUser(u *User) {
	db.Save(u)
}

func GetUser(email string) *User {
	u := User{}
	db.Where(User{
		Email: email,
	}).First(&u)
	return &u
}

func GetUserDataOneDayOnServer(email string, day string, region string, index int) *UserData {
	var ud UserData
	db.FirstOrInit(&ud, UserData{Email: email,
		Date:   day,
		Region: region,
		Index:  index})
	return &ud
}

func SaveUserData(ud *UserData) {
	db.Save(ud)
}

func SumUserData() {
	db.Exec("CALL sum_user_data();")
}
