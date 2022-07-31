package util

import (
	"regexp"
	"strconv"
	"strings"
)

// QQNum 正则表达式获取QQ号
func QQNum(str string) int64 {

	matched := regexp.MustCompile("[1-9][0-9]{4,}")

	submatch := matched.FindStringSubmatch(str)

	parseInt, err := strconv.ParseInt(submatch[0], 10, 64)
	if err != nil {
		return 0
	}

	return parseInt
}

// SetuIdentifyKeywords 判断setu命令关键字
func SetuIdentifyKeywords(command string) bool {
	commandKeywords := strings.Contains(command, "/来一张")

	if commandKeywords {
		return true
	}
	return false
}

// BanIdentifyKeywords 判断ban命令关键字
func BanIdentifyKeywords(command string) bool {
	commandKeywords := strings.Contains(command, "/ban")

	if commandKeywords {
		return true
	}
	return false
}

// UnBanIdentifyKeywords 判断unban命令关键字
func UnBanIdentifyKeywords(command string) bool {
	commandKeywords := strings.Contains(command, "/unban")

	if commandKeywords {
		return true
	}
	return false
}

// FindPermissionIdentifyKeywords 判断权限命令关键字
func FindPermissionIdentifyKeywords(command string) bool {
	commandKeywords := strings.Contains(command, "/我的权限")

	if commandKeywords {
		return true
	}
	return false
}
