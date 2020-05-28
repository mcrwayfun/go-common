package main

import (
	"fmt"
	"strings"
)

func main() {
	// 切成两部分[1 2,3,4]
	splitN := strings.SplitN("1,2,3,4", ",", 2)
	fmt.Println(splitN)

	// 全部切[1 2 3 4]
	splitN = strings.SplitN("1,2,3,4", ",", -1)
	fmt.Println(splitN)

	// n大于sep在s中出现的次数,也会全切
	splitN = strings.SplitN("1,2,3,4", ",", 10)
	fmt.Println(splitN)

	// n为0时,返回nil
	splitN = strings.SplitN("1,2,3,4", ",", 0)
	fmt.Println(splitN)
}
