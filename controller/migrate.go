package controller

import (
	"fmt"
	"gin-gorm-demo/database"
	"gin-gorm-demo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createTable(c *gin.Context) {

	fmt.Println("打印c", c)
	// 判断是否有passages表，如果没有创建表，如果有删除表
	hasTable := database.MYSQLDB.HasTable(&models.Passages{})
	fmt.Println("has table", hasTable)
	if hasTable {
		if err := database.MYSQLDB.DropTable(&models.Passages{}).Error; err != nil {
			fmt.Println("删除表失败", err)
		}
	} else {
		if err := database.MYSQLDB.CreateTable(&models.Passages{}).Error; err != nil {
			fmt.Println("创建passages表失败", err)
		}
	}
}

func DeleteTable(c *gin.Context) {
	hasTable := database.MYSQLDB.HasTable("passages")
	if hasTable {
		if err := database.MYSQLDB.DropTable("passages").Error; err != nil {
			fmt.Println("删除passage", err)
		}
	}
}
