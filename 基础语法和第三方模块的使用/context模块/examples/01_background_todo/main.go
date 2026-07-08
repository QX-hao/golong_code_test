package main

import (
	"context"
	"fmt"
)

// context.Context 是 Go 标准库 context 包中的一个接口。
//
// 这个接口有 4 个方法：
//   - Deadline() (time.Time, bool)
//     返回 context 的截止时间，以及是否设置了截止时间。
//     第二个返回值为 false，表示没有截止时间。
//   - Done() <-chan struct{}
//     返回一个只读 channel。
//     当 context 被取消或超时时，这个 channel 会被关闭。
//     对于 Background() 和 TODO() 这种永远不会主动结束的 context，Done() 返回 nil。
//   - Err() error
//     返回 context 结束的原因。
//     如果还没有被取消或超时，返回 nil。
//     如果被手动取消，通常返回 context.Canceled。
//     如果因为超时结束，通常返回 context.DeadlineExceeded。
//   - Value(key any) any
//     根据 key 读取 context 中保存的值。
//     如果没有对应的值，返回 nil。
func printContextState(ctx context.Context, name string) {
	// Deadline 返回两个值：
	// 第一个值是截止时间，这里暂时不用，所以用 _ 忽略。
	// 第二个值 hasDeadline 表示是否真的设置了截止时间。
	_, hasDeadline := ctx.Deadline()

	fmt.Println("context:", name)
	fmt.Println("  has deadline:", hasDeadline)

	// Err 返回当前 context 是否已经结束，以及结束原因。
	// Background 和 TODO 默认都没有结束，所以这里一般打印 <nil>。
	fmt.Println("  err:", ctx.Err())

	// Done 返回用于接收取消/超时通知的 channel。
	// Background 和 TODO 不会主动取消，也没有超时，所以 Done() 返回 nil。
	fmt.Println("  done channel is nil:", ctx.Done() == nil)
}

func queryUser(ctx context.Context, userID int64) {
	printContextState(ctx, "queryUser")
	fmt.Println("query user:", userID)
}

func main() {
	// Background 返回一个空的根 context，常作为整个调用链的起点。
	background := context.Background()

	// TODO 也返回一个空 context，通常表示这里暂时还不知道该传什么 context。
	todo := context.TODO()

	printContextState(background, "Background")
	printContextState(todo, "TODO")

	queryUser(background, 1001)
}
