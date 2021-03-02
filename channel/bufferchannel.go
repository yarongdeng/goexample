package main

import "fmt"

func main() {
	// 定义带缓冲区 channel
	var chn chan int = make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			chn <- i
		}
		close(chn)
	}()

	for num := range chn {
		fmt.Println(num)
	}
}
