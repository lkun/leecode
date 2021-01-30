package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)

	for _, n := range nums {
		m[n] = true
	}
	result := []int{}

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

func findDisappearedNumbers1(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		if v > 0 && nums[v-1] > 0 { //只有nums[v-1]>0才需置反
			nums[v-1] = -nums[v-1]
		}
		if v < 0 && nums[-v-1] > 0 { //同上
			nums[-v-1] = -nums[-v-1]
		}
	}
	for k, v := range nums {
		if v > 0 {
			result = append(result, k+1) //缺少这个元素
		}
	}
	return result
}
