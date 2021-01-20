# leecode
记录自己的 leecode 算法（基于 go 实现）

### 数组
* 1_两数和

```go
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num
		if j, ok := m[pair]; ok && i != j {
			return []int{j, i} // 注意返回值顺序，向后遍历 nums，所以 i 在 j 后
		}
		m[num] = i
	}
	return nil
}
```

* 485_最大连续1的个数
```go
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
```

* 495_提莫攻击
```go
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
```

* 414_第三大的数
```go
func thirdMax(nums []int) int {
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
```
### 链表


### 字符串