package main

import (
	"context"
	"fmt"
)

type contextKey string

const (
	requestIDKey contextKey = "request_id"
	traceIDKey   contextKey = "trace_id"
)

func logInfo(ctx context.Context, message string) {
	requestID, _ := ctx.Value(requestIDKey).(string)
	traceID, _ := ctx.Value(traceIDKey).(string)

	fmt.Printf("[request_id=%s trace_id=%s] %s\n", requestID, traceID, message)
}

func saveOrder(ctx context.Context, orderID string) {
	logInfo(ctx, "save order "+orderID)
}

func createOrder(ctx context.Context, orderID string) {
	logInfo(ctx, "create order "+orderID)
	saveOrder(ctx, orderID)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-1001")
	ctx = context.WithValue(ctx, traceIDKey, "trace-abc")

	createOrder(ctx, "order-9")
}
