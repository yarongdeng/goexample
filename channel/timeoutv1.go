package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task(result chan<- int64) {
	// 随机休眠
	internal := rand.Intn(10)
	fmt.Println("sleep:", internal)
	time.Sleep(time.Duration(internal) * time.Second)
	result <- time.Now().Unix()
}

func main() {
	rand.Seed(time.Now().Unix())
	var result chan int64 = make(chan int64)
	var timeout chan int = make(chan int)
	fmt.Println(time.Now())
	go task(result)

	go func() {
		time.Sleep(3 * time.Second)
		close(timeout)
	}()

	// 对管道进行操作，只要有一个成功就执行
	select {
	case r := <-result:
		fmt.Println("sucess", r)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
	fmt.Println(time.Now())
}
