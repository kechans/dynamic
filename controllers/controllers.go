package controllers

import "github.com/gin-gonic/gin"

func LoginEndpoint() string {
	return "str"
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/getuser", GetUser)
	}
	//router.GET("/getUser", GetUser)

	v2 := router.Group("/v2")
	{
		v2.GET("/getuser", GetUser)
	}
	return router
}
