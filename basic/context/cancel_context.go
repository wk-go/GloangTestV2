package main

// 测试context的cancel操作， 所有接受同一个上下文的协程都会响应cancel()

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	var num = make(chan int)

	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 所有的协程都会响应cancel操作
	for i := 0; i < 10; i++ {
		go sub(cancelCtx, i, num)
	}

	go func() {
		for i := 0; i < 1000000; i++ {
			num <- i
		}
	}()

	// 让子弹飞一会儿
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
		time.Sleep(3 * time.Second)
		ch <- struct{}{}
	}()
	<-ch
}

func sub(ctx context.Context, id int, num chan int) {
	for {
		select {
		case n := <-num:
			fmt.Printf("sub[%02d]: %d\n", id, n)
		case <-ctx.Done():
			fmt.Printf("sub[%02d] %s\n", id, ctx.Err())
			return
		}
	}
}
