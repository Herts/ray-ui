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

func AllUserDataOnServer(region string, index int, startDate, endDate string) (ds []*models.UserData) {
	s, _, req := MakeRequest(region, index, "/api/user/listData")
	if s.ApiKey == "" {
		return
	}
	resp := response{Data: &ds}
	req.Query(fmt.Sprintf("startDate=%s&endDate=%s", startDate, endDate)).
		EndStruct(&resp)
	for _, d := range ds {
		d.Region = region
		d.Index = index
	}
	return
}

func MakeRequest(region string, index int, path string) (*models.RemoteServer, url.URL, *gorequest.SuperAgent) {
	s := RemoteServer(region, index)
	u := url.URL{
		Scheme: "https",
		Host:   s.Host,
		Path:   path,
	}
	req := gorequest.New().Get(u.String()).Set("Authorization", s.ApiKey)
	return s, u, req
}

func UpdateAllUserDataOnServer(region string, index int, startDate, endDate string) {
	ds := AllUserDataOnServer(region, index, startDate, endDate)
	for _, d := range ds {
		data := models.GetUserDataOneDayOnServer(d.Email, d.Date, d.Region, d.Index)
		d.Model = data.Model
		models.SaveUserData(d)
	}
}

func AllUserOnServer(region string, index int) (userServers []*models.UserServer) {
	_, _, req := MakeRequest(region, index, "/api/user/list")
	resp := response{Data: &userServers}
	req.EndStruct(&resp)
	for _, us := range userServers {
		us.Region = region
		us.Index = index
	}
	return
}

func RetrieveAllUserOnServer(region string, index int) {
	userServers := AllUserOnServer(region, index)
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
		RetrieveAllUserOnServer(s.Region, s.Index)
	}
}

func UpdateAllUserDataOnAllServers(startDate, endDate string) {
	servers := models.GetAllServers()
	for _, s := range servers {
		UpdateAllUserDataOnServer(s.Region, s.Index, startDate, endDate)
	}
}
