package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Ban 禁言
func Ban(gid int64, uid int64, duiTime int64) {

	// 结构体
	msg := struct {
		Uid     int64 `json:"user_id"`
		Gid     int64 `json:"group_id"`
		DuiTime int64 `json:"duration"`
	}{
		Uid:     uid,
		Gid:     gid,
		DuiTime: duiTime,
	}

	payload, _ := json.Marshal(&msg)

	_, err := http.Post("http://localhost:5700/set_group_ban", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
}

// UnBan 解除禁言
func UnBan(gid int64, uid int64) {
	// 结构体
	msg := struct {
		Uid     int64 `json:"user_id"`
		Gid     int64 `json:"group_id"`
		DuiTime int64 `json:"duration"`
	}{
		Uid:     uid,
		Gid:     gid,
		DuiTime: 0,
	}

	payload, _ := json.Marshal(&msg)

	_, err := http.Post("http://localhost:5700/set_group_ban", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
}
