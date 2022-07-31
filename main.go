package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go_cqhttp_demo/api"
	"go_cqhttp_demo/util"
	"io/ioutil"
	"strconv"
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
			fmt.Println("QQ群号：" + strconv.FormatInt(gid, 10))

			// 发送者信息
			sender := gjson.Get(string(rawData), "sender").String()
			fmt.Println("发送者信息：" + sender)

			// QQ
			uid := gjson.Get(string(rawData), "user_id").Int()
			fmt.Println("QQ：" + strconv.FormatInt(uid, 10))

			// 消息类型
			subType := gjson.Get(string(rawData), "message_type").String()
			fmt.Println("消息类型：" + subType)

			// 权限群组
			role := gjson.Get(sender, "role").String()
			fmt.Println("权限群组：" + role)

			if message == "1" {
				api.SendMessage(subType, uid, gid, "0", false)
			}

			if util.SetuIdentifyKeywords(message) {
				msg := "请稍后。。。"
				api.SendMessage(subType, uid, gid, msg, false)

				c.JSON(200, gin.H{
					"reply":     "[CQ:image,file=" + api.GetSetu() + ",id=40000]",
					"at_sender": true,
				})
			}

			if util.BanIdentifyKeywords(message) {
				api.Ban(gid, util.QQNum(message), 100)
				c.JSON(200, gin.H{
					"reply":     "已禁言",
					"at_sender": true,
				})
			}

			if util.UnBanIdentifyKeywords(message) {
				api.UnBan(gid, util.QQNum(message))
				c.JSON(200, gin.H{
					"reply":     "已经解除禁言",
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
