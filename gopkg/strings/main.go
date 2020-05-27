package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.ContainsAny("abc","ab"))
	fmt.Println(strings.ContainsAny("abc","z"))

	fmt.Println(strings.IndexAny("abc","ab"))
	fmt.Println(strings.IndexAny("abc","z"))
	fmt.Println(strings.IndexAny("abc","97"))

	fmt.Println(strings.ContainsRune("geeksforgeeks", 97))
	fmt.Println(strings.ContainsRune("a", 97))

	fmt.Println(strings.EqualFold("abc","ABC"))
	fmt.Println(strings.EqualFold("abc","ABC"))
}
