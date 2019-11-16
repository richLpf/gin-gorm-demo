package controller

import (
	"gin-gorm-demo/conf"
	DB "gin-gorm-demo/database"
	"gin-gorm-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	user := models.Users{
		Sex:      1,
		Username: "lvpengfei",
		Password: "123434",
		Remark:   "",
		Mail:     "peng@163.com",
		Phone:    "182211212",
		MoneyNum: 0,
	}
	if err := DB.MYSQLDB.Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success"})
}
