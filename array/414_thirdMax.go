package main

import "fmt"

func main() {
	//fmt.Println(thirdMax([]int{1, 2, 3, 4, 5, 6, 7})) // 5
	//fmt.Println(thirdMax([]int{1, 2}))                // 2
	//fmt.Println(thirdMax([]int{1, 2, 2}))             // 2
	//fmt.Println(thirdMax([]int{2, 2, 3, 1}))          // 1
	fmt.Println(thirdMax1([]int{2, 5, 6, 3})) // 3
	//fmt.Println(thirdMax([]int{2, 5, 7, 6}))          // 6
}

func thirdMax1(nums []int) int {
	minNum := ^int(^uint(0) >> 1)
	one := nums[0]
	// 数组只有两个值的情况找最大值
	if len(nums) <= 2 {
		for _, v := range nums {
			if v > one {
				one = v
			}
		}
		return one
	}
	two := minNum
	three := minNum

	for i := 1; i < len(nums); i++ {
		now := nums[i]
		// 当前值比第一大的值要大，那么需要将原来第二大的值换成原来第一大的值，原来第三大的值换成原来第二大的值，原来第一大的值要换成当前值
		if now > one {
			three = two
			two = one
			one = now
		} else if now < one && now > two {// 当前值要比原第一大的值小比原第二大的值大，那么原第一大的值不变，原第二大的值换成当前值，原第三大的值换成原第二大的值
			three = two
			two = now
		} else if now < two && now > three {// 当前值比原第一大和第二大值值要小比原第三大的值大，那么原第一和第二大的值不变，原第三大的值换成当前值
			three = now
		}
	}

	if three == minNum {
		return one
	} else {
		return three
	}
}

func thirdMax(nums []int) int {
	n := len(nums)
	minNum := ^int(^uint(0) >> 1)
	one := nums[0]
	two := minNum
	three := minNum

	for i := 1; i < n; i++ {
		now := nums[i]
		// 如果当前值等于第一、第二、第三的值，跳过不处理（重复）
		if now == one || now == two || now == three {
			continue
		}
		// 当前值比第一大的值要大，那么需要将原来第二大的值换成原来第一大的值，原来第三大的值换成原来第二大的值，原来第一大的值要换成当前值
		if now > one {
			three = two
			two = one
			one = now
		} else if now < one && now > two { // 当前值要比原第一大的值小比原第二大的值大，那么原第一大的值不变，原第二大的值换成当前值，原第三大的值换成原第二大的值
			three = two
			two = now
		} else if now < two && now > three { // 当前值比原第一大和第二大值值要小比原第三大的值大，那么原第一和第二大的值不变，原第三大的值换成当前值
			three = now
		}
	}
	// 如果不存在第三大的值，那么返回原第一大的值
	if three == 0 {
		return one
	}
	return three
}
