package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// 数组是值，将一个数组赋予另外一个数组时，将会复制所有元素
	a := [3]int{1, 2, 3}
	b := [3]int{}
	b = a
	fmt.Println(b)

	fmt.Printf("%p\n", &a)
	// 若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针
	a = passArray(a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%v\n", a)

	m := make(map[string]Person, 0)
	p := Person{"gavin"}
	m["hello"] = p
	fmt.Printf("%p\n", &m)
	passMap(m)
	fmt.Printf("%v\n", m)

	var t interface{}
	t = m
	switch t := t.(type) {
	case map[string]Person:
		fmt.Printf("m is map[string]Person %T", t)
	default:
		fmt.Printf("un defined type %T", t)
	}
}

func passMap(m map[string]Person) {
	p := Person{"royce"}
	m["hello"] = p
	fmt.Printf("%p\n", m)
}

func passArray(c [3]int) [3]int {
	fmt.Printf("%p\n", &c)
	c[0] = 5
	return c
}
