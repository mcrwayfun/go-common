package main

import "fmt"

func main() {
	total := sum(1, 2)
	fmt.Printf("total=%d\n", total)

	// defer
	b()
	x := c()
	fmt.Printf("x=%d\n", x)
}

func sum(a, b int) (sum int) {
	fmt.Printf("sum=%d\n", sum)
	sum = a + b
	fmt.Printf("sum=%d\n", sum)
	return
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	un(trace("b"))
	fmt.Println("in b")
	a()
}

func c() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)
	x = 7
	return 9
}