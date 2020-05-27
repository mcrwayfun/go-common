package main

import (
	"fmt"
	"strings"
)

func main() {
	// 大小写
	fmt.Println(strings.EqualFold("abc", "Abc"))
	// 标题
	fmt.Println(strings.EqualFold("a b", "A B"))
	fmt.Println(strings.EqualFold("a b", "a B"))
	//output
	//true
	//true
	//true
}
