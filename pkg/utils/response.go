package utils

import (
	"github.com/gin-gonic/gin"
)

type ResultData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func Response(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, ResultData{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}
