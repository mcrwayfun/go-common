package main

import (
	"fmt"
	"strings"
)

func main() {
	// output:length:4, [1, 2, 3, 4]
	split := strings.SplitAfter("1,2,3,4", ",")
	fmt.Printf("length:%d, %v\n", len(split), split)

	// 若sep为空格,则会将s分成每一个unicode码值一个字符串
	// output:length:7, [1 , 2 , 3 , 4]
	split = strings.SplitAfter("1,2,3,4", "")
	fmt.Printf("length:%d, %v", len(split), split)
}
