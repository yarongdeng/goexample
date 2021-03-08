package main

import (
	"fmt"
	"time"
)

// 演示chan 定义 读写，以及阻塞

func main() {
	// 无缓冲区管道定义
	ch := make(chan int)

	go func() {
		fmt.Println("go 写入管道： 1")
		ch <- 1 //将1写入管道
		fmt.Println("go 写入管道： 2", time.Now())
		ch <- 2 //将2写入管道，管道第二个值写入要等待第一个值读出
		fmt.Println("go 管道写入结束", time.Now())
	}()

	fmt.Println("channel begin")
	fmt.Println("channel 读取管道第一个值")
	num := <-ch //如果未读取到数据会阻塞
	fmt.Println("管道值输出打印：", num)
	// time.Sleep(10 * time.Second)

	fmt.Println("管道第一个值读取完成后，sleep 5s 后，再次读取管道值，为演示阻塞")
	time.Sleep(5 * time.Second)

	fmt.Println("channel 读取管道第二个值")
	num = <-ch
	fmt.Println("管道值输出打印：", num)
	time.Sleep(1 * time.Second)


	fmt.Println("----------------------------------------------------------------------")

	//遍历，需要同时关闭管道，否则无法遍历，因为range长度无法控制
	go func() {
		ch <- 1
		fmt.Println("写入管道：1",time.Now())
		ch <- 2
		fmt.Println("写入管道：2",time.Now())
		ch <- 3
		fmt.Println("写入管道：3",time.Now())
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num,time.Now())
	}

	// 关闭管道，管道关闭不能写，但是可以读
	// close(ch)
	go func() {
		fmt.Println("close before")
		time.Sleep(3 * time.Second)
		close(ch)
		fmt.Println("close after")
	}()
	// 判断ch存在
	v, ok := <-ch
	fmt.Println(v, ok)

}
