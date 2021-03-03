package main

import (

	"fmt"
	"sync"
	"time"
)

type Context1 struct {
	lock sync.Mutex
	C chan int
	Values map[string]interface{}
	child []*Context1
}
type cancelCtx struct {
	Context1
	mu sync.Mutex
	done chan struct{}
}

func  (c *cancelCtx)Done()<-chan struct{}  {

	if c.done==nil {
		c.done=make(chan struct{})
	}
	d:=c.done

	return d
}
func WithCancel(parent Context1) (*cancelCtx) {
	c:=cancelCtx{Context1:parent}
	return &c
}

func cancel(c cancelCtx) {

		close(c.done)



}


func worker( ctx *cancelCtx, t *time.Ticker) {
	// Ticker是一个计时器，每过一段时间就向管道Ticker.C中发送一个Time结构体
	go func() {
		defer fmt.Println("worker exit")

		// Using stop channel explicit exit
		for {
			select {
			case <-ctx.Done():
				// ctx.Done()是一个方法，返回一个chan struct{}类型的管道
				// 当主函数中调用了cancel()方法，就会关闭这个管道
				// 按照我们前面讲到，这时会从管道拿到一个零值，用于通知结束
				fmt.Println("recv stop signal")
				return
			case <-t.C:
				fmt.Println("is working")
				// do something
			}
		}
	}()
	return
}

func main() {
	ticker := time.NewTicker(time.Second) // 创建一个计时器
	Background:=Context1{}
	// 创建一个context，Background()方法返回一个空的context，表示最上层
	// ctx是一个Context类型，cancel是一个func() 类型的函数
	ctx:=WithCancel(Background)

	go worker(ctx, ticker)

	time.Sleep(3 * time.Second)
	go cancel(*ctx)
	time.Sleep(50 * time.Millisecond) // 给一定的时间打印信息
}