package controller

import (
	"fmt"
	"gin-gorm-demo/common"
	"gin-gorm-demo/conf"
	"gin-gorm-demo/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通过将字段变成key,value形式，省掉大量的重复代码
type ReqRegionInfo struct {
	Query string   `json:"query"`
	Value []string `json:"value"`
}
type Regions struct {
	Id       int    `gorm:"AUTO_INCCREMENT,primary_key" json:"id"`
	RegionCn string `json:"region_cn"`
	RegionEn string `json:"region_en"`
	RegionId int    `json:"region_id"`
	ZoneCn   string `json:"zone_cn"`
	ZoneEn   string `json:"zone_en"`
	ZoneId   int    `json:"zone_id"`
	Address  string `json:"address"`
}

func GetRegionInfo(c *gin.Context) {
	hastable := common.HasTable("regions")
	if !hastable {
		if err := database.MYSQLDB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Regions{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_Fail, "message": err})
			return
		}
		fmt.Println("create databse ok")
	}
	var reqRegionInfo ReqRegionInfo
	resRegionInfo := make([]Regions, 0)
	if err := c.ShouldBindJSON(&reqRegionInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_Fail, "message": err})
		return
	}
	query := reqRegionInfo.Query + " IN (?)"
	if err := database.MYSQLDB.Debug().Table("regions").Where(query, reqRegionInfo.Value).Find(&resRegionInfo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_OK, "message": "success", "data": resRegionInfo})
}
