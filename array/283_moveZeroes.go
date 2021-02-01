package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	left := 0
	right := 0
	n := len(nums)
	for right < n {
		if nums[right] != 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
		}
		right++
	}
}
