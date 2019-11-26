package http

import (
	"fmt"
	"gin-gorm-demo/controller"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gin-gorm-demo/docs"

	"github.com/gin-gonic/gin"
	"gin-gorm-demo/common"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(MiddlewareMongo())
	//router.Use(gin.Logger())
	router.Use(gin.Recovery())

	web := router.Group("/web/passage")
	web.GET("/list", controller.GetPassageList)
	web.POST("/add", controller.AddPassage)
	web.GET("/detail/:id", controller.PassageDetail)
	web.POST("/edit/:id", controller.EditPassage)
	web.POST("/del/:id", controller.DelPassage)

	admin := router.Group("/web/admin")
	admin.POST("/signup", controller.SignUp)
	admin.POST("/signin", controller.SignIn)

	public := router.Group("/public")
	public.POST("/send", common.SendToMails)

	api := router.Group("/v1")
	api.GET("/region/list", controller.GetRegionInfo)

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
