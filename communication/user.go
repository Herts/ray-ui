package communication

import (
	"fmt"
	"github.com/Herts/ray-ui/models"
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

func AllUserDataOnServer(region string, index int, startDate, endDate, host, apiKey string) (ds []*models.UserData) {
	_, req := MakeRequest(host, apiKey, "/api/user/listData")
	resp := response{Data: &ds}
	req.Query(fmt.Sprintf("startDate=%s&endDate=%s", startDate, endDate)).
		EndStruct(&resp)
	for _, d := range ds {
		d.Region = region
		d.Index = index
	}
	return
}

func MakeRequest(host string, apiKey string, path string) (url.URL, *gorequest.SuperAgent) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	req := gorequest.New().Get(u.String()).Set("Authorization", apiKey)
	return u, req
}

func UpdateAllUserDataOnServer(region string, index int, startDate, endDate, host, apiKey string) {
	ds := AllUserDataOnServer(region, index, startDate, endDate, host, apiKey)
	for _, d := range ds {
		data := models.GetUserDataOneDayOnServer(d.Email, d.Date, d.Region, d.Index)
		d.Model = data.Model
		models.SaveUserData(d)
	}
}

func AllUserOnServer(region string, index int, host, apiKey string) (userServers []*models.UserServer) {
	_, req := MakeRequest(host, apiKey, "/api/user/list")
	resp := response{Data: &userServers}
	req.EndStruct(&resp)
	for _, us := range userServers {
		us.Region = region
		us.Index = index
	}
	return
}

func RetrieveAllUserOnServer(region string, index int, host, apiKey string) {
	userServers := AllUserOnServer(region, index, host, apiKey)
	for _, d := range userServers {
		data := models.GetUserOnServer(d.Email, d.Region, d.Index)
		d.Model = data.Model
		d.UpDataConsumed = data.UpDataConsumed
		d.DownDataConsumed = data.DownDataConsumed
		models.SaveUserServer(d)
	}
}

func RetrieveAllUserOnAllServers() {
	servers := models.GetAllServers()
	for _, s := range servers {
		RetrieveAllUserOnServer(s.Region, s.Index, s.Host, s.ApiKey)
	}
}

func UpdateAllUserDataOnAllServers(startDate, endDate string) {
	servers := models.GetAllServers()
	for _, s := range servers {
		UpdateAllUserDataOnServer(s.Region, s.Index, startDate, endDate, s.Host, s.ApiKey)
	}
}
