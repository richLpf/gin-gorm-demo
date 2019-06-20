package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MYSQLDB *gorm.DB

func main() {
	var err error
	// gorm
	MYSQLDB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/lpf?charset=utf8&parseTime=True&loc=Local")
	MYSQLDB.SingularTable(true)
	if err != nil {
		fmt.Println("connection err")
	} else {
		fmt.Println("connection succedssed")
	}
	defer MYSQLDB.Close()
	// router
	router := InitRouter()
	router.Run("0.0.0.0:3000")
}
