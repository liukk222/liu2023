package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 创建超时上下文
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	// 模拟耗时的后台任务
	go simulateTask(ctx)

	// 等待并处理请求
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintln(w, "Request processed successfully")
	case <-ctx.Done():
		fmt.Fprintln(w, "Request timed out")
	}
}

func simulateTask(ctx context.Context) {
	select {
	case <-time.After(8 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled")
	}
}
