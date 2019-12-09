package common

import (
	"errors"
	"fmt"
	"gin-gorm-demo/conf"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// companyid
var corpid string = "test"

// appid
var corpsecret string = "test"

// 获取token
var token_url string = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"

// 应用id
var agentid int = 1000000

// 创建群聊
var create_chat_url string = "https://qyapi.weixin.qq.com/cgi-bin/appchat/create"

// 发送企业微信
var send_info_url string = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send"

// 推送个人企业微信
var send_member_url string = "https://qyapi.weixin.qq.com/cgi-bin/message/send"

// 企业微信登录授权获取信息
var qy_chat_info string = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"

// 获取企业微信成员信息
var qy_chat_user_get string = "https://qyapi.weixin.qq.com/cgi-bin/user/get"

// 缓存access_token
var glo_access_token string = ""
var glo_expires_in string = ""

/***************************获取access_token**********************************/
func GetQyToken() (res interface{}, expires_temp_time string, err error) {
	url := token_url + "?corpid=" + corpid + "&corpsecret=" + corpsecret
	res, err = Get(url)
	if err != nil {
		return res, expires_temp_time, err
	}
	// 计算失效时间
	now := time.Now()
	mm, err := time.ParseDuration("100m")
	if err != nil {
		return res, expires_temp_time, err
	}
	cur_expires_in := now.Add(mm)
	expires_time := time.Date(cur_expires_in.Year(), cur_expires_in.Month(), cur_expires_in.Day(), cur_expires_in.Hour(), cur_expires_in.Minute(), cur_expires_in.Second(), cur_expires_in.Nanosecond(), time.Local)
	expires_temp_time = expires_time.Format("2006-01-02 15:04:05")
	return res, expires_temp_time, err
}

// 1、判断是否有access,如果没有，获取access, 计算过期时间点
// 2、如果有access, 判断是否过期，如果没有过期，使用，过期，重新执行第一步
// 3、缓存access
func GetAccessToken() (result string, err error) {
	var res interface{}
	if glo_access_token == "" {
		res, glo_expires_in, err = GetQyToken()
		if err != nil {
			return "", err
		}
		response := res.(map[string]interface{})
		if response["errcode"].(float64) != 0 {
			return "", errors.New(response["errmsg"].(string))
		}
		glo_access_token = response["access_token"].(string)
		fmt.Println("第一次获取", glo_access_token)
	} else {
		loc, _ := time.LoadLocation("PRC")
		stringToTime, _ := time.ParseInLocation("2006-01-02 15:04:05", glo_expires_in, loc)
		beforeOrAfter := stringToTime.After(time.Now())
		// 过期
		if beforeOrAfter == false {
			res, glo_expires_in, err = GetQyToken()
			if err != nil {
				return "", err
			}
			response := res.(map[string]interface{})
			if response["errcode"].(int64) != 0 {
				return "", errors.New(response["errmsg"].(string))
			}
			glo_access_token = response["access_token"].(string)
			fmt.Println("过期，重新获取", res, glo_expires_in, err)
		} else {
			fmt.Println("不过期，可以继续使用")
		}
	}
	result = glo_access_token
	return result, err
}

/********************************转发各种类型消息**********************************/
func SendMessage(c *gin.Context) {
	//获取参数
	sendtype := c.Param("sendtype")
	fmt.Println("sendType", sendtype)
	accessToken, err := GetAccessToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//转发
	sendUrl := ""
	switch sendtype {
	case "appchat":
		sendUrl = send_info_url + "?access_token=" + accessToken
	case "message":
		sendUrl = send_member_url + "?access_token=" + accessToken
	case "createchat":
		sendUrl = create_chat_url + "?access_token=" + accessToken
	default:
		c.JSON(http.StatusOK, gin.H{"message": "Got wrong sendType,Expected in appchat,message,createchat", "code": conf.Ret_NotFound})
		return
	}
	var sm interface{}
	if err := c.ShouldBindJSON(&sm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": conf.Ret_Fail, "error": err.Error()})
		return
	}
	ret, err := Post(sendUrl, sm, "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": conf.Ret_Fail, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ret)
}

// 通过企业微信登录页面授权信息，返回用户信息
type QyPageInfo struct {
	Code string `json:"code"`
}

func GetQyUserId(access_token string, code string) (userid string, err error) {
	url := qy_chat_info + "?access_token=" + access_token + "&code=" + code
	result, err := Get(url)
	if err != nil {
		return "", err
	}
	response := result.(map[string]interface{})
	userid = response["UserId"].(string)
	return userid, err
}

func GetQyUserInfo(c *gin.Context) {
	var req QyPageInfo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "params error " + err.Error()})
		return
	}
	accessToken, err := GetAccessToken()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "get token " + err.Error()})
		return
	}
	code := req.Code
	userid, err := GetQyUserId(accessToken, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "get user " + err.Error()})
		return
	}
	fmt.Println("userid", userid)
	getUserInfoUrl := qy_chat_user_get + "?access_token=" + accessToken + "&userid=" + userid

	res, err := Get(getUserInfoUrl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": conf.Ret_Fail, "message": "get userinfo error" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
