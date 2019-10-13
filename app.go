package main

import (
	"fmt"
	"gin-gorm-demo/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var database.MYSQLDB *gorm.DB

func main() {
	var err error
	// gorm
	database.MYSQLDB, err = gorm.Open("MYSQLDB", "root:123456@tcp(127.0.0.1:3306)/lpf?charset=utf8&parseTime=True&loc=Local")
	database.MYSQLDB.SingularTable(true)
	if err != nil {
		fmt.Println("connection err")
	} else {
		fmt.Println("connection succedssed")
	}
	defer database.MYSQLDB.Close()
	// router
	router := InitRouter()
	router.Run("0.0.0.0:3000")
}
