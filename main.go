package main

import (
	"github.com/dukeduffff/home_service/cmd"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 创建一个默认的Gin路由器
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableQuote:  true,
	})
	engine := gin.Default()
	cmd.Route(engine)
	err := engine.Run(":3010")
	if err != nil {
		log.Errorf("xray start error=%s", err)
		return
	}
}
