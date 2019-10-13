package controller

import (
	"crypto/md5"
	"fmt"
	"gin-gorm-demo/conf"
	"gin-gorm-demo/database"
	"gin-gorm-demo/models"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*注册*/
func SignUp(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	user.Password = PassMd5(user.Password)
	fmt.Println("user.Password", user.Password)
	var total int
	database.MYSQLDB.Where("username = ?", user.Username).First(&user).Count(&total)
	fmt.Println("total", total)
	if total == 0 {
		// 密码需要加密
		if err := database.MYSQLDB.Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "account is exist"})
		return
	}
}

/*获取用户信息*/
func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	if err := database.MYSQLDB.Where("deleted = ? AND id = ?", 0, id).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success", "data": user})
}

/*登录*/
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录
func SignIn(c *gin.Context) {
	var userinfo UserInfo
	var total int
	c.BindJSON(&userinfo)
	header := c.Request
	fmt.Println("获取header信息", header)

	// 判断字段不能为空
	if userinfo.Username == "" || userinfo.Password == "" {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "params is required"})
		return
	}
	var user models.Users
	if err := database.MYSQLDB.Where("username = ?", userinfo.Username).Find(&user).Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
		return
	}
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "no user"})
		return
	}
	fmt.Println("password", userinfo.Password)
	password := PassMd5(userinfo.Password)
	fmt.Println("加密", userinfo.Password)
	if err := database.MYSQLDB.Where("username = ? AND password = ?", userinfo.Username, password).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "password error"})
		return
	} else {
		// 匹配成功，后插入登录表中，生成token,将用户信息记录 // 登录成功返回，token

		if err := database.MYSQLDB.Exec("update login_token set valid = ? WHERE username = ?", 2, userinfo.Username).Error; err != nil {
			fmt.Println("将之前的登录信息失效")
		}
		var login_token models.LoginToken
		login_token.Username = userinfo.Username
		login_token.Token = TokenMd5()
		login_token.ExpireAt = time.Now().Unix() + 36000
		login_token.Valid = 1
		if err := database.MYSQLDB.Create(&login_token).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
			return
		}
		if err := database.MYSQLDB.Where("username = ? AND deleted = ?", userinfo.Username, 0).First(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_OK, "message": "success", "data": user, "token": login_token.Token})
	}
}

// md5加密
func PassMd5(str string) (md5str string) {
	data := []byte(str)
	fmt.Println("data", data)
	has := md5.Sum(data)
	fmt.Println("has", has)
	md5str = fmt.Sprintf("%x", has)
	return md5str
}

// md5获取token
func TokenMd5() string {
	curtime := time.Now().Unix()
	fmt.Println("curtime", curtime)
	h := md5.New()
	fmt.Println("h-->", h)
	fmt.Println("strconv.FormatInt(curtime, 10)-->", strconv.FormatInt(curtime, 10))
	io.WriteString(h, strconv.FormatInt(curtime, 10))

	fmt.Println("h-->", h)

	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("token--->", token)
	return token
}

// 判断token是否有效
func TokenIsValid(str string) bool {
	var login_token models.LoginToken
	if err := database.MYSQLDB.Where("token = ? AND valid = ", str, 1).First(&login_token).Error; err != nil {
		fmt.Println("err", err.Error())
	}
	curtime := time.Now().Unix()
	if curtime > login_token.ExpireAt {
		// 过期
		return false
	}
	/*	expire_date := time.Now().Unix() + 36000
		if err := database.MYSQLDB.Exec("update login_token set expire_at = ? Where token = ? AND valid = 1", expire_date, str).Error; err != nil {
			fmt.Println("err", err.Error())
		}
	*/return true

}
