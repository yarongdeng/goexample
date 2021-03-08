package main

import "fmt"

func add(a, b int, s ...int) int {
	sum := a + b
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	// ... pack(打包) unpack(解包)
	fmt.Println(append(s1, s2...))
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
}
