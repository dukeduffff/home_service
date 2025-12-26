package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var keys = []string{"SrqyTnrhcV6FEwz8URLacX"}

func SendMessage(ctx *gin.Context) {
	title := ctx.Query("title")
	msg := ctx.Query("msg")
	badge := ctx.DefaultQuery("badge", "1")
	group := ctx.DefaultQuery("group", "")
	jsonMap := map[string]interface{}{
		"title": title,
		"body":  msg,
		"badge": badge,
		"group": group,
	}
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		log.Errorf("json marshal error=%s", err)
	}
	for _, key := range keys {
		if response, err := http.Post(
			fmt.Sprintf("https://api.day.app/%s", key), "application/json",
			bytes.NewBuffer(jsonBytes)); err != nil {
			log.Errorf("send message error=%s, resp=%v", err, response)
		}
	}

	response := gin.H{
		"message": "ok",
		"code":    0,
	}
	ctx.JSON(http.StatusOK, response)
}
