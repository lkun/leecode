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

* 628_三个数的最大乘积
```go
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
```

* 645_错误的集合
```go
方法一：使用 Map
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

方法二：使用额外空间
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

方法三：桶法
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
```
### 链表


### 字符串