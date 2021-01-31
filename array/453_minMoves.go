package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minMoves([]int{1, 2, 3}))
	//fmt.Println(minMoves([]int{1, 2, 3, 5}))
	fmt.Println(minMoves1([]int{1, 2, 3, 4}))
}

func minMoves1(nums []int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	count := 0
	for _, num := range nums {
		count = count + num - min
	}
	return count
}
func minMoves(nums []int) int {
	moves := 0
	min := math.MaxInt64
	for n := 0; n < len(nums); n++ {
		min = Min(min, nums[n])
	}

	for n := 0; n < len(nums); n++ {
		moves = moves + nums[n] - min
	}
	return moves
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
