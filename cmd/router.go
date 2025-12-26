package cmd

import (
	"github.com/dukeduffff/home_service/common"
	"github.com/dukeduffff/home_service/xray"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.Engine) {
	// 文件系统
	engine.StaticFS("/subscribe", http.Dir("./static"))
	// 订阅更新接口
	engine.GET("/:PathParam/add_config", xray.AddConfig)
	// 生成订阅文件
	engine.GET("/gen_config", xray.GenAllConfigs)
	// 发送信息
	engine.GET("/send_message", common.SendMessage)
}
