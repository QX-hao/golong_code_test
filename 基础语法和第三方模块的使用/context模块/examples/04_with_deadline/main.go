package main

import (
	"context"
	"fmt"
	"time"
)

func phase(ctx context.Context, name string, cost time.Duration) error {
	fmt.Println(name, "started")

	select {
	case <-time.After(cost):
		fmt.Println(name, "finished")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	deadline := time.Now().Add(800 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	if err := phase(ctx, "phase 1", 300*time.Millisecond); err != nil {
		fmt.Println("phase 1 error:", err)
		return
	}

	if err := phase(ctx, "phase 2", 700*time.Millisecond); err != nil {
		fmt.Println("phase 2 error:", err)
		return
	}

	fmt.Println("all phases finished")
}
