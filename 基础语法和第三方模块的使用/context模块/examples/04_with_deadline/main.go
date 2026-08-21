package main

import (
	"context"
	"fmt"
	"time"
)

// phase 用来模拟一个阶段任务。
//
// ctx context.Context：
//   - 外部传进来的 context，用来接收取消或超时信号。
//   - 这里两个 phase 共用同一个 ctx，所以它们共享同一个截止时间。
//
// name：阶段名称，只用于打印。
// cost：这个阶段模拟需要花费的时间。
//
// 返回值 error：
//   - 返回 nil，表示这个阶段在截止时间前完成。
//   - 返回 ctx.Err()，表示 context 已经取消或超时。
func phase(ctx context.Context, name string, cost time.Duration) error {
	fmt.Println(name, "started")

	// select 同时等待两个 channel：
	//   - time.After(cost)：表示这个阶段正常耗时结束。
	//   - ctx.Done()：表示 context 被取消或到达截止时间。
	select {
	case <-time.After(cost):
		// cost 时间先到，说明这个阶段正常完成。
		fmt.Println(name, "finished")
		return nil
	case <-ctx.Done():
		// ctx.Done() 先触发，说明 context 已经取消或超时。
		// 如果是 WithDeadline 到时间了，ctx.Err() 通常是 context.DeadlineExceeded。
		return ctx.Err()
	}
}

func main() {
	// WithDeadline 和 WithTimeout 都能做超时控制，区别在于参数不同：
	//
	//   - WithTimeout(parent, timeout)
	//     传入的是“一段时间”，例如 800*time.Millisecond。
	//     含义是：从创建 context 开始，最多再运行 800 毫秒。
	//
	//   - WithDeadline(parent, deadline)
	//     传入的是“具体截止时间点”，类型是 time.Time。
	//     含义是：到了 deadline 这个具体时间点，context 就超时。
	//
	// 这里先用 time.Now().Add(...) 算出“当前时间 800 毫秒之后”的具体时间点。
	deadline := time.Now().Add(800 * time.Millisecond)

	// WithDeadline 基于父 context 创建一个带截止时间的子 context。
	// 到达 deadline 后，ctx.Done() 会被关闭，ctx.Err() 会变成 context.DeadlineExceeded。
	//
	// 这段代码效果接近：
	//   context.WithTimeout(context.Background(), 800*time.Millisecond)
	//
	// 区别是 WithDeadline 直接接收一个 time.Time 截止时间点。
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	// 即使 ctx 会自动超时，也建议调用 cancel。
	// 如果任务提前结束，cancel 可以释放 context 内部的定时器资源。
	defer cancel()

	// phase 1 需要 300 毫秒。
	// 总截止时间是 800 毫秒，所以 phase 1 会先完成。
	if err := phase(ctx, "phase 1", 300*time.Millisecond); err != nil {
		fmt.Println("phase 1 error:", err)
		return
	}

	// phase 1 已经用掉大约 300 毫秒，此时距离 deadline 只剩大约 500 毫秒。
	// phase 2 需要 700 毫秒，所以 deadline 会先到，ctx.Done() 会触发。
	if err := phase(ctx, "phase 2", 700*time.Millisecond); err != nil {
		fmt.Println("phase 2 error:", err)
		return
	}

	// 只有两个阶段都在 deadline 前完成，才会执行到这里。
	fmt.Println("all phases finished")
}
