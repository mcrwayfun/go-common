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
3. 数组的大小是其类型的一部分。类型 `[10]int` 和 `[20]int` 是不同的

