## New 分配

Go提供了两种分配原句，分别是new和make，这两者是不同的

Let's talk about `new` first. It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not *initialize* the memory, it only *zeros* it. That is, `new(T)` allocates zeroed storage for a new item of type `T` and returns its address, a value of type `*T`. In Go terminology, it returns a pointer to a newly allocated zero value of type `T`

new是用来分配内存的内建函数，与其他语言中的同名函数不同的是。简单的理解上述这段话，new的作用是为类型申请一片内存空间，并返回指向这片内存的指针。

## 构造函数与复合字面

## make分配

内建函数 `make(T, args)` 的目的不同于 `new(T)`。它只用于创建切片、映射和信道，并返回类型为 `T`（而非 `*T`）的一个**已初始化** （而非**置零**）的值。 出现这种用差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化。

```go
var p *[]int = new([]int)       // 分配切片结构；*p == nil；基本没用
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 习惯用法：
v := make([]int, 100)
```

使用new来创建切片，其没有初始化，需要另外再初始化

## 数组

1. 数组是值，将一个数组赋予另外一个数组时，将会复制所有元素
2. 若将某个数组传入某个函数，它将收到该数组的一份副本而非指针
   1. 传递的是副本？为什么可以修改数组的值？
3. 数组的大小是其类型的一部分。类型 `[10]int` 和 `[20]int` 是不同的

## 切片

切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针

## 映射

映射是强大的内建数据结构，可以关联不同类型的值，即常说的key-value结构。

1. 键可以是任何**相等性操作符**支持的类型，如整数、浮点数、复数、字符串、指针、接口，结构或数组等。但是切片不能用作映射键。

2. 试图通过映射中不存在的键来取值，就会返回与该映射中项的类型对应的零值

3. 可以通过逗号ok的方式来判断键是否存在于映射中

   ```go
   seconds, ok := timeZone[tz];if ok{
     // exist
   }
   ```

4. 要删除映射中的某项，可使用内建函数 `delete`，它以映射及要被删除的键为实参。 即便对应的键不在该映射中，此操作也是安全的。

## append

1. 可以将值追加到切片上

   ```go
   var str []string
   str = append(str,"hello")
   var str1 []string
   // 追加切片
   str = append(str,str1...)
   ```


## 问题

1. map是值传递（拷贝），但是在func中修改值后也会影响到原本。

```go
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

// output
[1 2 3]
0xc00009c000
0xc00009c040
0xc00009c000
[5 2 3]
0xc000092020
0xc000074180
map[hello:{royce}]
m is map[string]Person map[string]main.Person
Process finished with exit code 0
```





