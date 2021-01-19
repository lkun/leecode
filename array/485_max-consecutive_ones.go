package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{}))                 // 0
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // 3
	fmt.Println(findMaxConsecutiveOnes1([]int{1, 1, 0, 1, 1, 1})) // 3
}

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	count, maxCount := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count = 1
			for j := i + 1; j < n; j++ {
				if nums[j] == 0 {
					break
				}
				count++
				i++
			}
			maxCount = max(count, maxCount)
		}
	}
	return maxCount
}

func findMaxConsecutiveOnes1(nums []int) int {
	count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count += 1
		} else {
			maxCount = max(count, maxCount)
			count = 0
		}
	}
	return max(count, maxCount)
}

func max(count int, count2 int) int {
	if count > count2 {
		return count
	}
	return count2
}
