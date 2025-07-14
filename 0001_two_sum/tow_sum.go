package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 数值 -> 下标

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i} // 找到就返回两个下标
		}
		numMap[num] = i // 没找到就记录当前数字和下标
	}
	return nil // 如果没有找到（题目说一定有答案）
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // 输出：[0 1]
}

