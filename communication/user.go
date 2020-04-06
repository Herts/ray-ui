package communication

import (
	"../models"
	"github.com/astaxie/beego/logs"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RemoteServer(region string, index int) *models.RemoteServer {
	s := &models.RemoteServer{Region: region, Index: index}
	s = models.GetServer(s)
	return s
}

func AllUserDataOnServer(region string, index int) (ds []*models.UserData) {
	s := RemoteServer(region, index)
	u := url.URL{
		Scheme: "https",
		Host:   s.Host,
		Path:   "/api/user/listData",
	}
	req := gorequest.New()
	resp := response{Data: &ds}
	req.Get(u.String()).Set("Authorization", s.ApiKey).EndStruct(&resp)
	for _, d := range ds {
		d.Region = region
		d.Index = index
	}
	logs.Debug(ds[0].UpDataConsumed)
	return
}
