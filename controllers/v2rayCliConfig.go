package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/Herts/ray-ui/models"
	"github.com/astaxie/beego"
)

type V2RayConfigController struct {
	beego.Controller
}

// @Param email path string true "User email"
// @Param quan query bool false "Query quantumult config"
// @Param token query string true "Access Token"
// @router /get/:email [get]
func (c *V2RayConfigController) GetUserServerV2rayConfig(email string, quan bool, token string) {
	u := models.GetUser(email)
	if u.AccessToken == "" || u.AccessToken != token {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Ctx.ResponseWriter.Write([]byte("token is invalid or empty"))
		return
	}
	configs := models.GetV2rayNConfigByEmail(email)
	var configUrls string
	for _, config := range configs {
		if quan {
			configUrls += fmt.Sprintf("%s\n", models.Base64EncodeQuantumult(config, "vmess"))
		} else {
			configUrls += fmt.Sprintf("%s\n", models.Base64EncodeV2rayN(config, "vmess"))
		}
	}

	configUrls = base64.StdEncoding.EncodeToString([]byte(configUrls))
	c.Ctx.ResponseWriter.Write([]byte(configUrls))
}

// @Param email path string true "User email"
// @Param quan query bool false "Query quantumult config"
// @router /debug/get/:email [get]
func (c *V2RayConfigController) GetUserServerV2rayConfigDebug(email string, quan bool) {
	configs := models.GetV2rayNConfigByEmail(email)
	var quantumultConfigs []string
	if quan {
		for _, config := range configs {
			configUrl := models.GetRawQuantumultConfigUrl(config, "vmess")
			quantumultConfigs = append(quantumultConfigs, configUrl)
		}
		c.Data["json"] = quantumultConfigs
		c.ServeJSON()
		return
	}
	c.Data["json"] = configs
	c.ServeJSON()
}
