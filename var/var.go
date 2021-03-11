package main

import "fmt"

// 全局变量
// 全局变量可以定义不使用
var i int = 1

func main() {
	// 局部变量
	// 局部变量不能定义不使用
	var j string = "j"
	// 省略类型，由值推到，并且可以同时多个赋值，类型可以不同
	var a, b = 1, ""
	// 短变量声明，只能在函数内部使用
	c := 0.12
	fmt.Println("hello,world!", j, a, b, c)
}
