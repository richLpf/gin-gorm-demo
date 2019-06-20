package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPassageList(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	fmt.Println("header", token)
	//isLogined := TokenIsValid(token)
	/*if !isLogined {
		c.JSON(http.StatusOK, gin.H{"code": Ret_Not_Login, "message": "请登录"})
		return
	}*/
	// 增加分页
	passage_list := make([]Passages, 0)
	if err := MYSQLDB.Where("is_deleted = ?", 0).Find(&passage_list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": Ret_Fail, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": Ret_OK, "message": "success", "data": passage_list})
}

func AddPassage(c *gin.Context) {
	var passage Passages
	c.BindJSON(&passage)
	// 需要验证下标题是否存在
	if err := MYSQLDB.Create(&passage).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}

func EditPassage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}

func PassageDetail(c *gin.Context) {
	id := c.Param("id")
	var passage Passages
	if err := MYSQLDB.Where("id = ?", id).Find(&passage).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": passage})
}
func DelPassage(c *gin.Context) {
	id := c.Param("id")
	if err := MYSQLDB.Exec("UPDATE passages SET is_deleted = 1 WHERE id IN (?)", id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": nil})
}
