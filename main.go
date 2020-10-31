package main

import (
	"crypto/tls"
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func main() {
	router := gin.Default()
	router.GET("/email/:sender/:title/:dou", func(c *gin.Context) {
		title := c.Param("title")
		sender := c.Param("sender")
		dou := c.Param("dou")
		msg := gomail.NewMessage()
		msg.SetHeader("To", sender)
		msg.SetHeader("Subject", title)
		msg.SetBody("text/plain", "您在专业研报终端-研报客的会员，还剩"+dou+"天将到期，豆子与推送到期将清零！请及时续费！网页yanbaoke.com 手机app各大商店搜索 研报客")
		msg.SetAddressHeader("From", "mail@sendfile.yanbaoke.net", "研报客-全球研报速达")
		d := gomail.NewPlainDialer("192.168.1.20", 657, "mail", "360tHKWznZCY4Fy3")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		if err := d.DialAndSend(msg); err != nil {
			fmt.Println(err)
		}

	})
	router.Run(":8000")
}
