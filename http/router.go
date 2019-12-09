package http

import (
	"fmt"
	"gin-gorm-demo/controller"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gin-gorm-demo/docs"

	"gin-gorm-demo/common"

	"github.com/gin-gonic/gin"
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
	// 推送邮件
	public.POST("/mail/send", common.SendToMails)
	// ufile上传服务
	public.POST("/upload", common.Upload)
	// 转发企业微信服务
	public.POST("/userinfo", common.GetQyUserInfo)
	public.POST("/forword/:sendtype", common.SendMessage)

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
