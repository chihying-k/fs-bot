package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_Main(t *testing.T) {
	str := "/ban [CQ:at,qq=3186519875]"
	matched := regexp.MustCompile("[1-9][0-9]{4,}")

	submatch := matched.FindStringSubmatch(str)

	print("结果：")
	fmt.Println(submatch[0])
}
