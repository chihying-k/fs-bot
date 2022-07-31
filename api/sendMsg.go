package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendMessage 发送消息
func SendMessage(messageType string, uid int64, gid int64, message string, auto bool) {
	msg := struct {
		MessageType string `json:"message_type"`
		Uid         int64  `json:"user_id"`
		Gid         int64  `json:"group_id"`
		Message     string `json:"message"`
		Auto        bool   `json:"auto_escape"`
	}{
		MessageType: messageType,
		Uid:         uid,
		Gid:         gid,
		Message:     message,
		Auto:        auto,
	}

	payload, _ := json.Marshal(&msg)

	_, err := http.Post("http://localhost:5700/send_msg", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
}

/*

上面那个通用接口，以下接口废弃（白写了（bushi

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(uid int64, gid int64, message string, auto bool) {
	msg := struct {
		Uid     int64  `json:"user_id"`
		Gid     int64  `json:"group_id"`
		Message string `json:"message"`
		Auto    bool   `json:"auto_escape"`
	}{
		Uid:     uid,
		Gid:     gid,
		Message: message,
		Auto:    auto,
	}

	payload, _ := json.Marshal(&msg)

	_, err := http.Post("http://localhost:5700/send_private_msg", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}

}

// SendGroupMsg 发送群聊消息
func SendGroupMsg(gid int64, message string, auto bool) {
	msg := struct {
		Gid     int64  `json:"group_id"`
		Message string `json:"message"`
		Auto    bool   `json:"auto_escape"`
	}{
		Gid:     gid,
		Message: message,
		Auto:    auto,
	}

	payload, _ := json.Marshal(&msg)

	_, err := http.Post("http://localhost:5700/send_group_msg", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
}

*/
