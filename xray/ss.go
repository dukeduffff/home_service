package xray

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

type SSConfig struct {
	EncryptionMethod string
	Password         string
	Ip               string
	Port             string
	Ps               string
}

var ssConfigs []*SSConfig

func (s *SSConfig) AddConfig(ctx *gin.Context) {
	encryption := ctx.Query("encryption")
	password := ctx.Query("password")
	ip := ctx.Query("ip")
	port := ctx.Query("port")
	ps := ctx.Query("ps")
	lock.Lock()
	defer lock.Unlock()
	ssConfigs = append(ssConfigs, &SSConfig{
		EncryptionMethod: encryption,
		Password:         password,
		Ip:               ip,
		Port:             port,
		Ps:               ps,
	})
	ctx.JSON(200, gin.H{
		"message": "ok",
		"status":  0,
	})
}

func (s *SSConfig) GenConfig() ([]string, error) {
	if len(ssConfigs) == 0 {
		return nil, errors.New("ssConfigs is empty")
	}
	configs := make([]string, 0)
	for _, config := range ssConfigs {
		configStr := strings.Join([]string{strings.Join([]string{config.EncryptionMethod, config.Password}, ":"), strings.Join([]string{config.Ip, config.Port}, ":")}, "@")
		base64Str := base64.StdEncoding.EncodeToString([]byte(configStr))
		configs = append(configs, "ss://"+base64Str+"#"+config.Ps)
	}
	return configs, nil
}

func (s *SSConfig) Clear() {
	lock.Lock()
	defer lock.Unlock()
	ssConfigs = []*SSConfig{}
}
