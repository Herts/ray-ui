package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type V2RayConfig struct {
	V       string `json:"v"`
	Name    string `json:"ps"`
	Address string `json:"add"`
	Port    int    `json:"port,string"`
	UserId  string `json:"id"`
	AlterId int    `json:"aid,string"`
	Network string `json:"net"`
	Type    string `json:"type"`
	Host    string `json:"host"`
	Path    string `json:"path"`
	TLS     string `json:"tls"`
	Group   string `json:"-"`
}

func GetV2rayNConfigByEmail(email string) (configs []*V2RayConfig) {
	uss := GetVUserServerByEmail(email)
	for _, us := range uss {
		if us.Enabled == false {
			continue
		}
		config := V2RayConfig{
			V:       "2",
			Name:    GetServerNickName(us),
			Address: us.ServerName,
			Port:    us.Port,
			UserId:  us.UserId,
			AlterId: us.AlterID,
			Network: us.StreamSetting,
			Type:    us.ConfsType,
			Host:    us.Host,
			Path:    us.Path,
			TLS:     us.Tls,
			Group:   us.Region,
		}
		configs = append(configs, &config)
	}
	return
}

func GetServerNickName(us *VUserServer) string {
	return fmt.Sprintf("%s|%s|%s|%d", us.UserNickName, us.ServerNickName, us.Region, us.Index)
}

func Base64EncodeV2rayN(config *V2RayConfig, protocol string) (configUrl string) {
	byteConfig, err := json.Marshal(config)
	if err != nil {
		logs.Error(err, "when marshalling", config)
	}
	configUrl = fmt.Sprintf("%s://%s", protocol, base64.StdEncoding.EncodeToString(byteConfig))
	return
}

func Base64EncodeQuantumult(config *V2RayConfig, protocol string) (configUrl string) {
	configUrl = GetRawQuantumultConfigUrl(config, protocol)
	configUrl = fmt.Sprintf("%s://%s", protocol, base64.StdEncoding.EncodeToString([]byte(configUrl)))
	return configUrl
}

func GetRawQuantumultConfigUrl(config *V2RayConfig, protocol string) string {
	tlsSetting := fmt.Sprintf("tls-host=%s, over-tls=true, certificate=1,", config.Host)
	configUrl := fmt.Sprintf(`%[1]s = %[2]s, %[3]s, %[4]d, %[5]s, "%[6]s", group=%[7]s, %[8]s obfs=%[9]s, obfs-header="Host: %[10]s"`,
		config.Name, protocol, config.Address, config.Port, "chacha20-ietf-poly1305", config.UserId, config.Group,
		tlsSetting, config.Network, config.Host)
	return configUrl
}
