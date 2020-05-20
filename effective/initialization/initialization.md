## 常量

1. Go中的常量就是不变量。它们在编译时创建，即便它们可能是函数中定义的局部变量。 常量只能是数字、字符（符文）、字符串或布尔值。由于编译时的限制， 定义它们的表达式必须也是可被编译器求值的常量表达式

2. 比如 `1<<3` 就是一个常量表达式，而`math.Sin(math.Pi/4)` 则不是，因为对 `math.Sin` 的函数调用在运行时才会发生

3. Go中枚举常量可以使用 iota 创建

   ```go
   type ByteSize float64
   
   const (
       // 通过赋予空白标识符来忽略第一个值
       _           = iota // ignore first value by assigning to blank identifier
       KB ByteSize = 1 << (10 * iota)
       MB
       GB
       TB
       PB
       EB
       ZB
       YB
   )
   ```

## 变量

变量的初始化与常量不同的是，其初始值可以是在运行时才被计算的一般表达式。

```go
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)
```

## init

![image-20200520231913071](https://i.loli.net/2020/05/20/tTVkB8FDKx9dCN7.png)

每个源文件都可以定义多个无参数的init方法，其执行顺序如上所述。

1. 如果当前文件包含其他包(import)，则会优先初始化其他包
2. 当前包的常量初始化
3. 当前包的变量初始化 
4. 当前包的init方法执行
5. main方法执行