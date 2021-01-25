package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maximumProduct2([]int{1, 2, 3}))
	fmt.Println(maximumProduct2([]int{1, 2, 3, 4}))
	//fmt.Println(maximumProduct2([]int{-1, -2, -3}))
	//fmt.Println(maximumProduct2([]int{-1, -2, -3, -4}))
	//fmt.Println(maximumProduct2([]int{-1, -2, -3, 4}))
}

// 排序法
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	//nums[0]*nums[1]*nums[n-1] 针对有正负的情况
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

// 线性扫描法
func maximumProduct2(nums []int) int {
	// 定义第一小和第二小
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
