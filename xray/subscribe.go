package xray

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dukeduffff/home_service/xray/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var vmessConfigs []*config.VmessConfig
var lock sync.Mutex

func appendVmessConfig(config *config.VmessConfig) {
	defer func() { lock.Unlock() }()
	lock.Lock()
	vmessConfigs = append(vmessConfigs, config)
}

func clearConfigs() {
	defer func() { lock.Unlock() }()
	lock.Lock()
	vmessConfigs = []*config.VmessConfig{}
}

func AddVmess(ctx *gin.Context) {
	ip := ctx.Query("ip")
	ps := ctx.Query("ps")
	port := ctx.Query("port")
	id := ctx.DefaultQuery("id", "1f8f05a1-1a29-4862-a91a-ecb2a4a5e272")
	net := ctx.DefaultQuery("net", "tcp")
	security := ctx.DefaultQuery("security", "none")
	appendVmessConfig(&config.VmessConfig{
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

func GenConfig(ctx *gin.Context) {
	if len(vmessConfigs) <= 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "无可生成的配置",
			"code":    1,
		})
	}
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
	// 写到文件中
	finalConfigStr := strings.Join(base64Strs, "\n")
	base64FinalConfigStr := base64.StdEncoding.EncodeToString([]byte(finalConfigStr))
	configPath := "./static/config.txt"
	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		log.Errorf("Failed to create directories: %v", err)
		return
	}
	os.WriteFile("./static/config.txt", []byte(base64FinalConfigStr), 0644)
	clearConfigs()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "生成成功",
		"code":    0,
	})
}
