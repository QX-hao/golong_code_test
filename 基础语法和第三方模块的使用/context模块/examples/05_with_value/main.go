package main

import (
	"context"
	"fmt"
)

// type 用来定义一个新类型。
// 这里定义了 contextKey 类型，它的底层类型是 string。
//
// context.WithValue 的 key 不推荐直接使用普通 string，
// 因为不同包里可能都使用 "request_id" 这样的字符串，容易发生 key 冲突。
// 使用自定义类型 contextKey 可以降低冲突风险。
//
// 注意：
//   - contextKey("request_id")
//   - string("request_id")
//
// 这两个值虽然内容一样，但类型不同，所以不是同一个 context key。
type contextKey string

const (
	// requestIDKey 和 traceIDKey 是专门用于 context.Value 查询的 key。
	requestIDKey contextKey = "request_id"
	traceIDKey   contextKey = "trace_id"
)

// logInfo 模拟一段日志打印逻辑。
// 它不直接接收 requestID、traceID 参数，而是从 ctx 里读取。
func logInfo(ctx context.Context, message string) {
	// ctx.Value(key) 会根据 key 从 context 中取值，返回类型是 any。
	//
	// 因为返回的是 any，所以这里用 .(string) 做类型断言：
	//   - 如果取出来的是 string，requestID 就拿到对应字符串。
	//   - 如果 key 不存在，或者值不是 string，第二个返回值 ok 会是 false。
	//
	// 这里用 _ 忽略了 ok，是为了让示例更简单。
	requestID, _ := ctx.Value(requestIDKey).(string)
	traceID, _ := ctx.Value(traceIDKey).(string)

	fmt.Printf("[request_id=%s trace_id=%s] %s\n", requestID, traceID, message)
}

// saveOrder 和 createOrder 都接收同一个 ctx。
// 只要调用链一路把 ctx 传下去，后面的函数就能继续读取其中的上下文信息。
func saveOrder(ctx context.Context, orderID string) {
	logInfo(ctx, "save order "+orderID)
}

func createOrder(ctx context.Context, orderID string) {
	logInfo(ctx, "create order "+orderID)
	saveOrder(ctx, orderID)
}

func main() {
	// Background 创建一个空的根 context。
	ctx := context.Background()

	// WithValue 不会修改原来的 ctx，而是基于原 ctx 创建一个新的 context。
	//
	// 入参：
	//   - ctx：父 context。
	//   - requestIDKey：用于以后查询这个值的 key。
	//   - "req-1001"：真正保存的 value。
	//
	// 返回值：
	//   - 一个新的 context.Context。
	//
	// 所以这里要重新赋值给 ctx，让后续代码继续使用带 request_id 的新 context。
	ctx = context.WithValue(ctx, requestIDKey, "req-1001")

	// 再基于上一步的新 ctx 包一层，保存 trace_id。
	//
	// 结构可以理解成：
	//   Background
	//     -> WithValue(requestIDKey, "req-1001")
	//     -> WithValue(traceIDKey, "trace-abc")
	//
	// Value(key) 查询时会从最外层开始找，找不到就往父 context 继续找。
	// 如果同一个 key 重复保存，Value(key) 会返回最近一次保存的值。
	ctx = context.WithValue(ctx, traceIDKey, "trace-abc")

	// createOrder 接收到的是最终这个 ctx。
	// 因此 createOrder 和它后面调用的 saveOrder 都能读取 request_id 和 trace_id。
	createOrder(ctx, "order-9")
}
