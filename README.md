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

### 链表


### 字符串