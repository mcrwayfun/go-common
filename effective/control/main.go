package main

import "fmt"

func main() {
	/**
	满足以下条件的,已经被声明过的变量v可以出现在``:=``的右边
	1. 本次声明与已声明的变量v处于同一作用域（若v已经在外层作用域中声明过了，则此次声明会再次创建一个新的变量）
	2. 在初始化中与其类型相应的值才能赋予 `v`，且在此次声明中至少另有一个变量是新声明的
	*/

	// 声明a和b
	a, b := words("hello", "world")
	fmt.Printf("a:%v,b:%v\n", &a, &b)
	// 本次声明a与已声明的变量a处于同一作用域
	// a是赋值,c是声明
	a, c := words("ni", "hao")
	fmt.Printf("a:%v,b:%v\n", &a, &c)

	func() {
		// 若v已经在外层作用域中声明过了，则此次声明会再次创建一个新的变量
		a, c := words("ni", "hao")
		fmt.Printf("a:%v,b:%v\n", &a, &c)
	}()

	// 报错,下面语句中不存在新的变量
	// a, b := words("hello", "world")
	// 报错,返回的值与a不是同一种类型
	// a, b := words2("hello", "world")
	var words []string
	// for index
	for i := 0; i < 5; i++ {
		words = append(words, fmt.Sprintf("s%d", i))
	}
	// for range collection
	for i, v := range words {
		fmt.Printf("i=%d,v=%v\n", i, v)
	}
Loop:
	switch {
	case true:
		fmt.Println("switch true")
		break Loop
		fallthrough
	case false:
		fmt.Println("switch false")
		fallthrough
	default:
		fmt.Println("switch default")
	}

	var z interface{}
	z = "hello"
	switch z := z.(type) {
	case int:
		fmt.Printf("integer %d\n", z)
	case string:
		fmt.Printf("string %s\n", z)
	default:
		fmt.Printf("unexpected type %T", z)
	}
}

func words(a, b string) (string, string) {
	return a, b
}

func words2(a, b string) (int, string) {
	return 1, b
}
