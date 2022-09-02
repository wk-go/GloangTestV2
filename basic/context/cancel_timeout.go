package main

// 测试context的timeout操作， 所有接受同一个上下文的协程都会自动退出

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	var num = make(chan int)

	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// 所有的协程都会响应超时状态
	for i := 0; i < 10; i++ {
		go subTimeout(timeoutCtx, i, num)
	}

	go func() {
		for i := 0; i < 100000; i++ {
			num <- i
		}
	}()

	// 让子弹再飞一会儿
	ch := make(chan struct{})
	go func() {
		time.Sleep(4 * time.Second)
		ch <- struct{}{}
	}()
	<-ch
}

func subTimeout(ctx context.Context, id int, num chan int) {
	for {
		select {
		case n := <-num:
			fmt.Printf("subTimeout[%02d]: %d\n", id, n)
		case <-ctx.Done():
			fmt.Printf("subTimeout[%02d] %s\n", id, ctx.Err())
			return
		}
	}
}
