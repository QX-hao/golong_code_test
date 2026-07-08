package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context, cost time.Duration) error {
	select {
	case <-time.After(cost):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func run(name string, timeout time.Duration, cost time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("%s: timeout=%v cost=%v\n", name, timeout, cost)

	if err := slowOperation(ctx, cost); err != nil {
		fmt.Println("  failed:", err)
		return
	}

	fmt.Println("  success")
}

func main() {
	run("fast call", time.Second, 300*time.Millisecond)
	run("slow call", 700*time.Millisecond, 1500*time.Millisecond)
}
