package main

import "fmt"

func main() {
	total := sum(1, 2)
	fmt.Printf("total=%d\n", total)
}

func sum(a, b int) (sum int) {
	fmt.Printf("sum=%d\n", sum)
	sum = a + b
	fmt.Printf("sum=%d\n", sum)
	return
}
