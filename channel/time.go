package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	// 返回值为管道，三秒钟之后向管道中写入当前时间
	fmt.Println(<-time.After(3 * time.Second))

	for now := range time.Tick(3 * time.Second) {
		fmt.Printf("%T,%v\n", now, now)
	}
}
