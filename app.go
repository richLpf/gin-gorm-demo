package main

import (
	"fmt"
	"gin-gorm-demo/database"
	"gin-gorm-demo/http"
	"gin-gorm-demo/models"
)

//_ "github.com/jinzhu/gorm/dialects/mysql"

func main() {
	//var err error

	database.InitMysql()
	defer database.MYSQLDB.Close()

	models.AutoMigrateDB()

	//mysql_conf := con.Mysql
	//connect_sql := mysql_conf.Username + ":" + mysql_conf.Password + "@tcp(" + mysql_conf.Addr + ":" + mysql_conf.Port + ")/" + mysql_conf.Database + "?"
	//database.MYSQLDB, err = gorm.Open("mysql", connect_sql+"charset=utf8&parseTime=True&loc=Local")
	//database.MYSQLDB.SingularTable(true) // User表表明默认为users,  如果设置了这一句，创建的表为user, 而不是用复数

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//return "prefix_" + defaultTableName
	//}

	//if err != nil {
	//fmt.Println("connection err")
	//} else {
	//switch con.Env {
	//case "development":
	//fmt.Println("current environment is development")
	//case "production":
	//fmt.Println("current environment is production")
	//case "test":
	//fmt.Println("current environment is test")
	//}
	//}
	//defer database.MYSQLDB.Close()
	//models.InitDB()

	var con database.BuildConf
	con.GetConf()

	fmt.Println("con", con)
	router := http.InitRouter()
	router.Run("0.0.0.0:" + con.Listen)
}
