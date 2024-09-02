package cmd

import (
	"github.com/dukeduffff/home_service/xray"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.Engine) {
	// 文件系统
	engine.StaticFS("/subscribe", http.Dir("./static"))
	// 订阅更新接口
	engine.GET("/add_vmess", xray.AddVmess)
	// 生成订阅文件
	engine.GET("/gen_config", xray.GenConfig)
}
