package xray

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type VmessConfig struct {
	Add      string `json:"add"`
	Id       string `json:"id"`
	Port     string `json:"port"`
	Ps       string `json:"ps"`
	Security string `json:"security"`
	Net      string `json:"tcp"`
	Sni      string `json:"sni"`
	V        string `json:"v"`
	Fp       string `json:"fp"`
	Type     string `json:"type"`
	Aid      string `json:"aid"`
	Host     string `json:"host"`
	Tls      string `json:"tls"`
}

var vmessConfigs []*VmessConfig

func appendVmessConfig(config *VmessConfig) {
	defer func() { lock.Unlock() }()
	lock.Lock()
	vmessConfigs = append(vmessConfigs, config)
}

func clearConfigs() {
	defer func() { lock.Unlock() }()
	lock.Lock()
	vmessConfigs = []*VmessConfig{}
}

func (v *VmessConfig) AddConfig(ctx *gin.Context) {
	ip := ctx.Query("ip")
	ps := ctx.Query("ps")
	port := ctx.Query("port")
	id := ctx.DefaultQuery("id", "1f8f05a1-1a29-4862-a91a-ecb2a4a5e272")
	net := ctx.DefaultQuery("net", "tcp")
	security := ctx.DefaultQuery("security", "none")
	appendVmessConfig(&VmessConfig{
		Add:      ip,
		Ps:       ps,
		Port:     port,
		Id:       id,
		Security: security,
		Net:      net,
		V:        "2",
		Fp:       "chrome",
		Sni:      "",
		Aid:      "0",
	})
	response := gin.H{
		"message": "ok",
		"code":    0,
	}
	ctx.JSON(http.StatusOK, response)
}

func (v *VmessConfig) GenConfig() ([]string, error) {
	var configBytes [][]byte
	for _, c := range vmessConfigs {
		bytes, err := json.Marshal(c)
		if err != nil {
			log.Errorf("json gen error=%s", err)
			continue
		}
		configBytes = append(configBytes, bytes)
	}
	// 生成配置字符串
	var base64Strs []string
	for _, cs := range configBytes {
		configStr := base64.StdEncoding.EncodeToString(cs)
		base64Strs = append(base64Strs, fmt.Sprintf("vmess://%s", configStr))
	}
	return base64Strs, nil
}

func (v *VmessConfig) Clear() {
	clearConfigs()
}
