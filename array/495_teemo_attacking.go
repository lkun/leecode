package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2)) // 4
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2)) // 3
	fmt.Println(findPoisonedDuration([]int{1, 1}, 2)) // 2
	fmt.Println(findPoisonedDuration([]int{1, 3}, 2)) // 4
	fmt.Println(findPoisonedDuration([]int{1, 5}, 2)) // 4
}

func findPoisonedDuration(nums []int, duration int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	total := 0
	for i := 0; i < n-1; i++ {
		diff := nums[i+1] - nums[i]
		total += min(diff, duration)
	}
	return total + duration
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
