package main

import (
	"dynamic/controllers"
	"dynamic/models/dao"
	"dynamic/pkg/log"
	"dynamic/pkg/setting"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func init() {
	setting.InitConfig()
	log.Setup()
	dao.InitMysqlConfig()
	dao.InitRedisConfig()
}
func main() {
	gin.SetMode("debug")
	router := controllers.InitRouter()
	router.Run(":8090")
}
