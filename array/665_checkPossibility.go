package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{3,4,2,3}))
}

func checkPossibility(nums []int) bool {
	n := len(nums)
	var cnt int
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			cnt++
			if cnt == 1 && (i == 0 || i == n-2 || (i > 0 && nums[i-1] <= nums[i+1]) || (i < n-2 && nums[i] <= nums[i+2])) {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
