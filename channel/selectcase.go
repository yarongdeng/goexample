package main

import (
	"fmt"
	"time"
)

func main() {
	chaneelA := make(chan int)
	channelB := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		channelB <- 1
	}()

	// 多和for {} 一起使用
	select {
	case v, ok := <-chaneelA:
		fmt.Println("a", v, ok)
	case v, ok := <-channelB:
		fmt.Println("b", v, ok)
	default:
		fmt.Println("default")
	}
}
