package controllers

import (
	"dynamic/models/dao"
	"dynamic/pkg/utils"
	"dynamic/service"
	"github.com/gin-gonic/gin"
	//"github.com/go-ini/ini"
	"net/http"
)

func GetUser(c *gin.Context) {
	res := service.GetUserdata()
	//for key, value := range res {
	//	log.Printf("%v : %v", key, value)
	//}

	dao.Set("list", res, 20)
	data := make(map[string]interface{})
	data["lists"] = res
	utils.Response(c, http.StatusOK, 0, "", data)
}
