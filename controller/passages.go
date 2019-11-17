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
// @Param title body string true  "文章标题" "test"
// @Param author body string  true  "作者"
// @Param tag body  string  true  "标签"
// @Param look body int  true  "点击数量"
// @Param category body string  true  "类别"
// @Param description body  string  true  "描述"
// @Param content body string  true  "内容"
// @Param img_link body string  true  "图片连接"
// @Success 200 {object} models.Passages
// @Router /web/passage/add [post]
func AddPassage(c *gin.Context) {
	var passage models.Passages
	c.BindJSON(&passage)
	fmt.Println("passage", passage)
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
