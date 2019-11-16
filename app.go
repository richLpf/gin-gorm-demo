package main

import (
	"gin-gorm-demo/database"
	"gin-gorm-demo/http"
	"gin-gorm-demo/models"
)

func main() {
	// 数据库初始化
	database.InitMysql()
	// main函数执行结束后，关闭数据库连接，这里为了方便调用，我直接在main函数中关闭
	// 如果需要每次调用的收连接数据库，可以将该方法定义接口，在函数中调用，关闭
	defer database.MYSQLDB.Close()
	// 数据库迁移
	models.AutoMigrateDB()
	// 获取配置文件的监听端口
	var con database.BuildConf
	con.GetConf()
	// 初始化router
	router := http.InitRouter()
	router.Run("0.0.0.0:" + con.Listen)
}
