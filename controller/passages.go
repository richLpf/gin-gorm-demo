package controller

import (
	"fmt"
	"gin-gorm-demo/conf"
	DB "gin-gorm-demo/database"
	"gin-gorm-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPassageList(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	fmt.Println("header", token)
	//isLogined := TokenIsValid(token)
	/*if !isLogined {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Not_Login, "message": "请登录"})
		return
	}*/
	// 增加分页
	passage_list := make([]models.Passages, 0)
	if err := DB.MYSQLDB.Find(&passage_list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success", "data": passage_list})
}

func AddPassage(c *gin.Context) {
	var passage models.Passages
	c.BindJSON(&passage)
	// 需要验证下标题是否存在
	if err := DB.MYSQLDB.Create(&passage).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}

func EditPassage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}

func PassageDetail(c *gin.Context) {
	id := c.Param("id")
	var passage models.Passages
	if err := DB.MYSQLDB.Where("id = ?", id).Find(&passage).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": passage})
}
func DelPassage(c *gin.Context) {
	id := c.Param("id")
	if err := DB.MYSQLDB.Exec("UPDATE passages WHERE id IN (?)", id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}
