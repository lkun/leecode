package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findErrorNums([]int{1, 1}))
	//fmt.Println(findErrorNums2([]int{1, 2, 2, 4}))
	//fmt.Println(findErrorNums2([]int{4, 6, 2, 4, 1, 3}))
}

// 使用 Map
func findErrorNums(nums []int) []int {
	m := make(map[int]int, len(nums))
	dup := 0
	missing := 1
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
			dup = v
		} else {
			m[v] = 1
		}
	}
	// 因为数组中的元素是从1开始，故m中的key自然得从1开始
	for i := 1; i < len(nums) + 1; i++ {
		if _, ok := m[i]; !ok {
			missing = i
			break
		}
	}
	return []int{dup, missing}
}

// 使用额外空间
func findErrorNums1(nums []int) []int {
	dup := -1
	missing := 1

	for _, n := range nums {
		if nums[int(math.Abs(float64(n)))-1] < 0 {
			dup = int(math.Abs(float64(n)))
		} else {
			nums[int(math.Abs(float64(n)))-1] *= -1
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = i + 1
		}
	}
	return []int{dup, missing}
}

// 桶法
func findErrorNums2(nums []int) []int {
	result := make([]int, 2)
	temp := make([]int, len(nums) + 1)
	for _, n := range nums {
		temp[n]++
	}
	for i := 1; i < len(temp); i++ {
		if temp[i] == 1 {
			continue
		}
		if temp[i] == 2 {
			result[0] = i
		} else {
			result[1] = i
		}
	}
	return []int{result[0], result[1]}
}
