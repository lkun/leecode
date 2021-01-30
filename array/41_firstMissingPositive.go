package main

import "fmt"

func main() {
	//fmt.Println(firstMissingPositive([]int{3, 4, -1, 1})) // 2
	fmt.Println(firstMissingPositive([]int{3, 4, 6, 7})) // 2
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num - 1] = -abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
