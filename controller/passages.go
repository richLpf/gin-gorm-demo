package controller

import (
	"fmt"
	"gin-gorm-demo/conf"
	DB "gin-gorm-demo/database"
	"gin-gorm-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPassageList godoc
// @Summary 文章列表
// @Description 描述信息
// @Tags 类别
// @Accept json
// @Produce json
// @Param limit query string false  "Limit"
// @Param offset query string false  "Offset"
// @Success 200 {string} string "ok"
// @Router /web/passage/list [get]
func GetPassageList(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	fmt.Println("header", token)
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	//isLogined := TokenIsValid(token)
	/*if !isLogined {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Not_Login, "message": "请登录"})
		return
	}*/
	// 增加分页
	db := DB.MYSQLDB.Offset(offset).Limit(limit)
	passage_list := make([]models.Passages, 0)
	if err := db.Order("updated_at desc").Find(&passage_list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err})
		return
	}
	fmt.Println("passage_list", passage_list)
	c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success", "data": passage_list})
}

// AddPassage godoc
// @Summary 添加文章
// @Description 添加文章接口
// @Tags 类别
// @Accept json
// @Produce json
// @Param body body models.Passages true  "请求参数"
// @Success 200 {int} models.Passages.ID
// @Router /web/passage/add [post]
func AddPassage(c *gin.Context) {
	var passage models.Passages
	if err := c.BindJSON(&passage); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	// 需要验证下标题是否存在
	if err := DB.MYSQLDB.Create(&passage).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "id": passage.ID})
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
