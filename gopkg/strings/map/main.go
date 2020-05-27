package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

func main(){
	// will return aaaaaaaaaa
	minRune := func(rune) rune { return 'a' }
	m := strings.Map(minRune, tenRunes(unicode.MaxRune))
	fmt.Println(m)

	// 验证返回的是一个拷贝
	s := "a"
	m = strings.Map(minRune, s)
	fmt.Println(m)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Data !=
		(*reflect.StringHeader)(unsafe.Pointer(&m)).Data)

	// 可以不返回拷贝,取决于mapping函数
	identity := func(r rune) rune {
		return r
	}
	orig := "Input string that we expect not to be copied."
	m = strings.Map(identity, orig)
	fmt.Println(m)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&orig)).Data ==
		(*reflect.StringHeader)(unsafe.Pointer(&m)).Data)

	// 返回-1的字符将被丢弃，而不会被替换
	// 输出:abc123
	trimSpaces := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}
	m = strings.Map(trimSpaces, "   abc    123   ")
	fmt.Println(m)
}

func tenRunes(ch rune) string {
	r := make([]rune, 10)
	for i := range r {
		r[i] = ch
	}
	return string(r)
}