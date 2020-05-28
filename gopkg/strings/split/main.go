package main

import (
	"fmt"
	"strings"
)

func main(){
	split := strings.Split("1 2 3 4", " ")
	fmt.Println(split)

	// 若sep为空格,则会将s分成每一个unicode码值一个字符串
	split = strings.Split("1234", "")
	fmt.Println(split)
}
