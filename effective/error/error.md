## Panic

我们提供了内建的 `panic` 函数，它会产生一个运行时错误并终止程序 （但请继续看下一节）。该函数接受一个任意类型的实参（一般为字符串），并在程序终止时打印。 它还能表明发生了意料之外的事情，比如从无限循环中退出了。

## Recover

当 `panic` 被调用后（包括不明确的运行时错误，例如切片检索越界或类型断言失败）， 程序将立刻终止当前函数的执行，并开始回溯Go程的栈，运行任何被推迟的函数。 若回溯到达Go程栈的顶端，程序就会终止。不过我们可以用内建的 `recover` 函数来重新或来取回Go程的控制权限并使其恢复正常执行。

调用 `recover` 将停止回溯过程，并返回传入 `panic` 的实参。 由于在回溯时只有被推迟函数中的代码在运行，因此 `recover` 只能在被推迟的函数中才有效。

`recover` 的一个应用就是在服务器中终止失败的Go程而无需杀死其它正在执行的Go程。

```go
func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}
```

由于直接从被推迟函数中调用 `recover` 时不会返回 `nil`， 因此被推迟的代码能够调用本身使用了 `panic` 和 `recover` 的库函数而不会失败。例如在 `safelyDo` 中，被推迟的函数可能在调用 `recover` 前先调用记录函数，而该记录函数应当不受Panic状态的代码的影响。