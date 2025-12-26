package xray

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var lock sync.Mutex

type GenProcessor interface {
	AddConfig(context *gin.Context)
	GenConfig() ([]string, error)
	Clear()
}

var processorMap = map[string]GenProcessor{
	"vmess": &VmessConfig{},
	"ss":    &SSConfig{},
}

func GenAllConfigs(context *gin.Context) {
	configs := make([]string, 0)
	for _, processor := range processorMap {
		cs, err := processor.GenConfig()
		processor.Clear() // 清除缓存的配置
		if err != nil {
			log.Errorf("Failed to generate configs: %v", err)
			continue
		}
		configs = append(configs, cs...)
	}
	if len(configs) == 0 {
		log.Info("No configurations to generate")
		return
	}
	fileContent := strings.Join(configs, "\n")
	configPath := "./static/config.txt"
	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		log.Errorf("Failed to create directories: %v", err)
		return
	}
	os.WriteFile(configPath, []byte(fileContent), 0644)
	context.JSON(http.StatusOK, gin.H{
		"message": "生成成功",
		"code":    0,
	})
}

func AddConfig(ctx *gin.Context) {
	path := ctx.Param("PathParam")
	if processor, ok := processorMap[path]; ok {
		processor.AddConfig(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Processor not found",
		"status":  1,
	})
}
