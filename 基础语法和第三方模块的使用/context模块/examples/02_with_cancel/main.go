package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker 表示一个后台工作任务，通常会放到 goroutine 里执行。
//
// ctx context.Context：
//   - 用来接收外部传进来的取消信号。
//   - main 调用 cancel() 后，ctx.Done() 会被关闭，worker 就能知道该退出了。
//
// wg *sync.WaitGroup：
//   - *sync.WaitGroup 是指针类型，表示指向某个 WaitGroup 变量的地址。
//   - 这里要传指针，是为了让 main 和 worker 操作同一个 WaitGroup。
//   - worker 结束时调用 wg.Done()，main 里的 wg.Wait() 才能继续往下执行。
func worker(ctx context.Context, wg *sync.WaitGroup) {
	// defer 表示函数退出前执行。
	// 不管 worker 是正常结束，还是因为 context 被取消而 return，都会执行 wg.Done()。
	defer wg.Done()

	// NewTicker 创建一个定时器。
	// 这里表示每隔 300 毫秒，ticker.C 这个 channel 就会产生一个 time.Time 时间值。
	ticker := time.NewTicker(300 * time.Millisecond)

	// worker 退出时停止 ticker，避免定时器继续占用资源。
	defer ticker.Stop()

	for {
		// select 用来同时等待多个 channel。
		// 哪个 case 对应的 channel 先准备好，就执行哪个 case。
		select {
		case <-ctx.Done():
			// ctx.Done() 返回一个只读 channel。
			// main 调用 cancel() 后，这个 channel 会被关闭。
			// 这里的 <-ctx.Done() 表示等待取消信号，不关心取出来的空值。
			fmt.Println("worker stopped:", ctx.Err())
			return
		case t := <-ticker.C:
			// ticker.C 每 300 毫秒发送一次当前时间。
			// t 是 time.Time 类型，所以可以调用 Format 格式化输出。
			fmt.Println("working at", t.Format("15:04:05.000"))
		}
	}
}

func main() {
	// Background 是一个空的根 context。
	// WithCancel 基于这个父 context 创建一个新的、可以手动取消的子 context。
	// 返回值：
	//   - ctx：新的 context.Context，传给 worker 使用。
	//   - cancel：取消函数，调用后 ctx.Done() 会被关闭，ctx.Err() 会变成 context.Canceled。
	ctx, cancel := context.WithCancel(context.Background())

	// WaitGroup 用来等待 goroutine 执行结束。
	var wg sync.WaitGroup

	// Add(1) 表示要等待 1 个任务完成。
	// 后面 worker 结束时会调用 Done()，把任务数量减 1。
	wg.Add(1)

	// go worker(...) 表示启动一个新的 goroutine 执行 worker 函数。
	// &wg 是取地址，把 main 中这个 WaitGroup 的地址传给 worker。
	go worker(ctx, &wg)

	// main goroutine 暂停 950 毫秒。
	// worker 每 300 毫秒打印一次，所以这里大约会打印 3 次 working。
	time.Sleep(950 * time.Millisecond)
	fmt.Println("main calls cancel()")

	// 调用 cancel 后，ctx.Done() 会被关闭。
	// worker 的 select 会收到取消信号，然后打印 ctx.Err() 并退出。
	cancel()

	// 等 worker 调用 wg.Done() 后，Wait() 才会结束。
	// 这样可以确保 main 不会比 worker 更早退出。
	wg.Wait()
	fmt.Println("main exits")
}
