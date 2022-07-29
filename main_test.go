package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_Main(t *testing.T) {
	str := "/ban 1234565 60"
	matched, _ := regexp.MatchString("^.[a-n]+\\s+[0-9]*\\s+[0-9]*$", str)

	print("结果")
	fmt.Println(matched)
}
