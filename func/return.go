package main

import "fmt"

func bubble(nums []int) {
	for j := 0; j < len(nums)-1; j++ {
		// 将切片第一个数字，换到排序的位置，如果是最大就是最后一个，如果不是的，就是相应的位置
		for i := 0; i < len(nums)-j-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		fmt.Println(nums)
	}
}

func main() {
	// slice 切片是引用类型
	b := []int{5, 4, 3, 2, 1}
	bubble(b)
	// fmt.Println(b)
}
