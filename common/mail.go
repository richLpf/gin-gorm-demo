package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
	"gin-gorm-demo/conf"
	"strconv"
)

type MailInfo struct {
	SendUser    []string `json:"send_user"`
	CopyUser    []string `json:"copy_user"`
	Subject     string   `json:"subject"`
	HtmlContent string   `json:"html_content"`
}

func SendToMails(c *gin.Context) {
	var mail_info MailInfo
	if err := c.BindJSON(&mail_info); err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_Fail, "message": err})
		return
	}

	err := SendMails(mail_info.SendUser, mail_info.CopyUser, mail_info.Subject, mail_info.HtmlContent)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_Fail, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RetCode": conf.Ret_OK, "message": "success"})
}

// 发送邮件功能待定
func SendMails(send_user []string, copy_user []string, subject string, html_content string) (err error) {
	if len(send_user) == 0 && len(copy_user) == 0 {
		return
	}
	//定义服务信息
	mailConn := map[string]string{
		"user": "1234@163.com",
		"pass": "123456",
		"host": "smtp.qiye.163.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", send_user...)
	m.SetHeader("Cc", copy_user...)
	//m.SetAddressHeader("Cc", copy_user, "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html_content)
	//m.Attach("/home/Alex/lolcat.jpg")
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"]) 

	if err = d.DialAndSend(m); err != nil {
		panic(err)
	}
	return err
}
