package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go_cqhttp_demo/api"
	"go_cqhttp_demo/util"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		dataReader := c.Request.Body
		rawData, _ := ioutil.ReadAll(dataReader)
		postType := gjson.Get(string(rawData), "post_type").String() // 获取上报类型

		if postType == "message" {

			// 信息
			message := gjson.Get(string(rawData), "message").String()
			fmt.Println("信息：" + message)

			// QQ群号
			gid := gjson.Get(string(rawData), "group_id").Int()

			// 发送者信息
			sender := gjson.Get(string(rawData), "sender").String()
			fmt.Println("发送者信息：" + sender)

			// 权限群组
			role := gjson.Get(sender, "role").String()
			fmt.Println("权限群组：" + role)

			// 发送者QQ号
			// qq := gjson.Get(sender, "user_id").String()

			if message == "1" {
				api.SendGroupMsg(gid, "0", false)
			}

			if util.SetuIdentifyKeywords(message) {
				msg := "请稍后。。。"
				api.SendGroupMsg(gid, msg, false)

				c.JSON(200, gin.H{
					"reply":     "[CQ:image,file=" + api.GetSetu() + ",id=40000]",
					"at_sender": true,
				})
			}

			if util.BanIdentifyKeywords(message) {
				c.JSON(200, gin.H{
					"reply":     "您发送的是ban命令",
					"at_sender": true,
				})
			}

			if util.FindPermissionIdentifyKeywords(message) {
				c.JSON(200, gin.H{
					"reply":     "您的权限组：" + role,
					"at_sender": true,
				})
			}

		}

	})

	r.Run(":8081")
}
