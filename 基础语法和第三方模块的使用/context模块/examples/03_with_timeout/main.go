package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation 用来模拟一个耗时操作。
//
// ctx context.Context：
//   - 外部传进来的 context，用来接收取消或超时信号。
//   - 如果 ctx 超时了，ctx.Done() 会被关闭。
//
// cost time.Duration：
//   - 表示这个模拟操作需要花多久才能完成。
//   - 例如 300*time.Millisecond 表示 300 毫秒后完成。
//
// 返回值 error：
//   - 返回 nil，表示操作在超时前正常完成。
//   - 返回 ctx.Err()，表示操作还没完成，context 已经取消或超时。
func slowOperation(ctx context.Context, cost time.Duration) error {
	// select 用来同时等待多个 channel。
	// 这里同时等待“操作完成”和“context 超时/取消”。
	select {
	case <-time.After(cost):
		// time.After(cost) 返回一个只读 channel。
		// cost 时间到了以后，这个 channel 会收到一个 time.Time 值。
		// 这里不关心具体时间值，只用它表示“耗时操作完成了”。
		return nil
	case <-ctx.Done():
		// ctx.Done() 返回一个只读 channel。
		// WithTimeout 创建的 ctx 到达超时时间后，这个 channel 会被关闭。
		// 如果手动调用 cancel()，这个 channel 也会被关闭。
		// ctx.Err() 用来拿到结束原因，例如 context.DeadlineExceeded。
		return ctx.Err()
	}
}

// run 用来运行一次带超时控制的操作。
//
// name：本次调用的名字，只用于打印。
// timeout：给 context 设置的最长等待时间。
// cost：模拟操作真正需要花费的时间。
func run(name string, timeout time.Duration, cost time.Duration) {
	// WithTimeout 基于父 context 创建一个带超时时间的子 context。
	//
	// 入参：
	//   - context.Background()：父 context，作为调用链的起点。
	//   - timeout：超时时间，类型是 time.Duration。
	//
	// 返回值：
	//   - ctx：新的 context.Context。timeout 时间到了以后，ctx.Done() 会被关闭。
	//   - cancel：取消函数。调用它可以提前取消 ctx，并释放内部定时器资源。
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// 即使 ctx 会自动超时，也建议 defer cancel()。
	// 如果操作提前完成，cancel() 可以及时释放 WithTimeout 内部创建的定时器资源。
	defer cancel()

	fmt.Printf("%s: timeout=%v cost=%v\n", name, timeout, cost)

	// 如果 slowOperation 在 timeout 之前完成，err 为 nil。
	// 如果 timeout 先到，err 通常是 context.DeadlineExceeded。
	if err := slowOperation(ctx, cost); err != nil {
		fmt.Println("  failed:", err)
		return
	}

	fmt.Println("  success")
}

func main() {
	// timeout 是 1 秒，cost 是 300 毫秒。
	// 操作会先完成，所以会打印 success。
	run("fast call", time.Second, 300*time.Millisecond)

	// timeout 是 700 毫秒，cost 是 1500 毫秒。
	// context 会先超时，所以会打印 failed: context deadline exceeded。
	run("slow call", 700*time.Millisecond, 1500*time.Millisecond)
}
