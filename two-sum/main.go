package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("nums: %v\ntarget: %d\n", nums, target)
	fmt.Printf("result: %v\n", twoSum(nums, target))
}

// 两数之和 使用map结构来做算法复杂度是O(n),利用空间来换
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := m[diff]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
