package controller

import (
	"fmt"
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

// 查询产品列表,简单get查询， 时间查询

// 查询列表数据，包含（分页，排序desc倒叙， asc正序, 模糊查询， 批量查询，post）
// 列表查询一般按产品的尿性，会查很多东西，这里我们换成post查询
func QueryData(c *gin.Context) {
	passage := make([]models.Passages, 0)
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "20")
	db := DB.MYSQLDB.Offset(offset).Limit(limit)
	if err := db.Order("updated_at desc").Find(&passage).Offset(offset).Limit(limit).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success", "data": passage})
}

// 查询详情，各种参数查询
func QueryDetail(c *gin.Context) {
	var passage models.Passages
	query := ""
	if err := DB.MYSQLDB.Where(query).First(&passage).Error; err != nil {
		fmt.Println("err", err)
	}
}

// 连表查询,定义返回结构
type ResPassages struct {
	models.Passages
}

/*查询表中关联的用户id, 文章列表中关联用户id*/
/*func GetPassages(c *gin.Context) {
	DB.MYSQLDB.Table("pasages").Select("passages.*, user.name").joins("left join user on user.user_id = passages.author_id").Scan()
}*/

// 事务查询
/*业务复杂时，需要将*/

// 查询链

// 原生语句查询
