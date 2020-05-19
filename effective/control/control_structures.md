## if

1. if能够接初始化语句，因此可以用来设置局部变量

   ```go
   if err := file.Chmod(0664); err != nil {
   	log.Print(err)
   	return err
   }
   ```

2. if可用于控制流程的扭转，Go最通俗可见的便是 if err then return，即遇错则不再执行

   ```go
   f, err := os.Open(name)
   if err != nil {
   	return err
   }
   // 不再被执行
   codeUsing(f)
   ```

## 重新声明与再次赋值

```go
// 使用短变量声明f和err
f, err := os.Open(name)
// 又通过其他方式再次声明f和err
d, err := f.Stat()
```

1. 调用f.Stat，看上去似乎再次声明了d和err。尽管err已经在第一条语句中被声明了，但这是合理的。第二条语句中，它仅是被**重新赋值**

2. 在满足一下条件时，已经被声明过的变量v可以出现在``:=``的右边

   1. 本次声明与已声明的变量v处于同一作用域（若v已经在外层作用域中声明过了，则此次声明会再次创建一个新的变量）
   2. 在初始化中与其类型相应的值才能赋予 `v`，且在此次声明中至少另有一个变量是新声明的

3. Go中函数的形参和返回值虽然处于大括号之外，但它们的作用域与该函数体相同

   ```go
   // 示例中，a、b、c、sum处于同一作用域
   func add(a,b int)(sum int){
     c := a + b
     return c
   }
   ```

## for

1. GO不再支持do或while循环,只有一个更加通用的for，它有三种形式

   ```go
   // 初始化值的循环
   for init;condition;post{}
   // while循环
   for condition{}
   // for(;;)
   for {}
   ```

2. GO不支持逗号操作符，而++和--为语句而非表达式

## switch

1. switch 的表达式无须为常量或者整数，可以为任意表达式。case会自上而下的逐一进行求值直到匹配为止

2. 若switch后没有表达式，它将匹配true

   ```go
   switch  {
   	case true:
   		fmt.Println("switch true")
   	case false:
   		fmt.Println("switch false")
   	default:
   		fmt.Println("switch default")
   	}
   // output
   switch true
   ```

3. switch并不会自动下溯，即执行对应的case后会自动break。但是case可以通过逗号来列举相同的条件

   ```go
   func shouldEscape(c byte) bool {
   	switch c {
   	case ' ', '?', '&', '=', '#', '+', '%':
   		return true
   	}
   	return false
   }
   ```

4. 使用fallthrough可以强制执行后面的case代码

   ```go
   switch  {
   	case true:
   		fmt.Println("switch true")
   		fallthrough
   	case false:
   		fmt.Println("switch false")
   		fallthrough
   	default:
   		fmt.Println("switch default")
   	}
   // output
   switch true
   switch false
   switch default
   ```

5. 使用break直接跳出整个switch

   ```go
   Loop:
   	switch  {
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
   // output
   switch true
   ```

## 类型选择

1. `switch` 也可用于判断接口变量的动态类型。如 **类型选择** 通过圆括号中的关键字 `type` 使用类型断言语法

   ```go
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
   // output
   switch true
   string hello
   ```

   