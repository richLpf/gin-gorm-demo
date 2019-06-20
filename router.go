package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(MiddlewareMongo())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	web := router.Group("/web/passage")
	web.GET("/list", GetPassageList)
	web.POST("/add", AddPassage)
	web.GET("/detail/:id", PassageDetail)
	web.POST("/edit/:id", EditPassage)
	web.POST("/del/:id", DelPassage)

	admin := router.Group("/web/admin")
	admin.POST("/signup", SignUp)
	admin.POST("/signin", SignIn)

	return router
}

// 定义中间件
func MiddlewareMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
		//err := mongo.SetSession()
		method := c.Request.Method
		fmt.Println("method", method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json; charset=utf-8")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()

	}
}
