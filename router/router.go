package router

import (
	"ginblog/api/v1"
	"ginblog/config"
	

	"github.com/gin-gonic/gin"
)

func CreateRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	// 简单的路由组: v1
	router := r.Group("/v1")
	{
		router.POST("/user/add", v1.AddUser)
		router.DELETE("/user/delete/:id", v1.DeleteUser)
		router.POST("/user/update", v1.UpdateUser) // localhost:3000/v1/user/update?id=4 body&password=123456&username=csaj
		router.GET("/user/query", v1.QueryUser) // localhost:3000/v1/user/query?id=2
		router.GET("/user/queryall", v1.QueryUserList) // localhost:3000/v1/user/queryall?offset=0&size=10 偏移量 单页大小
	}

	r.Run(config.HttpPort)
}