package main

import (
	"fmt"
)

func main() {
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2, 1}))
}

func findShortestSubArray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := len(nums)               // 计数
	degree := 1                      // 数组的度
	degreeArray := make(map[int]int) // 数组中每个元素出现的频数
	degreeIndex := make(map[int]int) // 数组中频数第一次出现的下标
	for i, num := range nums {
		degreeArray[num]++
		if degreeArray[num] == 1 {
			degreeIndex[num] = i
		} else {
			l := i - degreeIndex[num] + 1
			if degree < degreeArray[num] || (degree == degreeArray[num] && count > l) {
				degree = degreeArray[num]
				count = l
			}
		}
	}
	if len(degreeArray) == len(nums) {
		return 1
	}
	return count
}

func findShortestSubArray1(nums []int) int {
	need := make(map[int]int)
	// 数字出现的最大次数
	most := 0
	for _, v := range nums {
		need[v]++
		if need[v] > most {
			most = need[v]
		}
	}
	left, right, l := 0, 0, len(nums)
	shortest := l
	supply := make(map[int]int)
	for right < l {
		curRight := nums[right]
		supply[curRight]++
		for supply[curRight] == most {
			curLeft := nums[left]
			supply[curLeft]--
			if supply[curRight] < most {
				shortest = min(shortest, right-left+1)
				break
			}
			left++
		}
		right++
	}
	return shortest
}

func min(a int, nums ...int) int {
	result := a
	for _, v := range nums {
		if v < result {
			result = v
		}
	}
	return result
}
