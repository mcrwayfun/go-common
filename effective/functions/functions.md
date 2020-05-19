## 多值返回

1. Go支持函数返回多值。这个特性与遇错即返回的特性完美结合

## 可命名结果形参

1. Go函数的返回值或结果 “形参“ 可被命名，并作为常规变量使用。一旦该函数开始执行，它们就会被初始化为对应类型的零值。若函数执行了一条不带实参的return语句，则结果行参的当前值将被返回。

   ```go
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
   // output
   sum=0
   sum=3
   total=3
   ```

   


