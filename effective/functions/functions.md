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

## Defer（Important)

Go的defer用于预设一个函数调用，该函数会在执行 defer 的函数返回之前立即执行。典型的就是解锁互斥和释放资源。

```go
// Contents 将文件的内容作为字符串返回。
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()  // f.Close 会在我们结束后运行。

	for {
		n, err := f.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err  // 我们在这里返回后，f 就会被关闭。
		}
	}
	return string(result), nil // 我们在这里返回后，f 就会被关闭。
}
```

值得注意的是，**defer一般在离打开最近的地方声明关闭**

The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the *defer* executes, not when the *call* executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions.

被推迟函数的实参（如果该函数为方法则还包含接受者）在推迟执行时就会求值，而不是在调用执行时才求值，这样无须担心变量值在函数执行时被改变。

被推迟的函数按照后进先出的顺序执行！

```go
func main() {
	// defer
	b()
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
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

// output
entering:b
in b
entering:a
in a
leaving:a
leaving:b
```

如何**如果该函数为方法则还包含接受者的**？

```go
func c() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)
	x = 7
	return 9
}

func main(){
  x := c()
	fmt.Printf("x=%d\n", x)
}

// output
in defer x as parameter: x = 0
in defer x after return: x = 9
x=9

```

可以分为三步来理解这句话：如果函数为方法且包含接受者，即示例中的defer function(n){}(x)

1. 会在声明defer时求值，而不会在执行defer时才求值。即x为0，因为此时x并未赋值
2. 推迟执行，此时会将 func(n int)(0) 推入栈中
3. 当遇到return或者panic时，会执行  func(n int)(0)
   1. ```fmt.Printf("in defer x as parameter: x = %d\n", n)``` 输出x=0，因为入栈时n=0
   2. ```fmt.Printf("in defer x after return: x = %d\n", x)``` 输出x=9，因为返回了9，将x赋值了。这还表明 defer func与整个方法是同一个作用域




