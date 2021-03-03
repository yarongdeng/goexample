package main

import "fmt"

// 只读管道
func in(chanel chan<- int) {
	chanel <- 1
	chanel <- 2
	chanel <- 3
}

// 只写管道
func out(channel <-chan int) {
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func main() {
	// 读写管道，可以当作只读或者只写使用
	// 只读管道不能写，只写管道不能读
	ch := make(chan int, 3)
	in(ch)
	out(ch)
}
