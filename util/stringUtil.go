package util

import "strings"

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

// FindPermissionIdentifyKeywords 判断权限命令关键字
func FindPermissionIdentifyKeywords(command string) bool {
	commandKeywords := strings.Contains(command, "/我的权限")

	if commandKeywords {
		return true
	}
	return false
}
