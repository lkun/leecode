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

* 697_数组的度
>  其实这道算法题，是求最大度（出现次数最多的数字），第一个出现的数字到最后一个出现的数字，中间相隔最少数字,
>eg: {1, 2, 2, 3, 1, 4, 2},出现数字最多的数字是2，出现了3次（其实就是度），那么从第一个2到最后一个2，
>组成的新的子数组为 {2，2，3，1，4，2}，长度为 6
>
>假如数组变为 {1, 2, 2, 3, 1, 4, 2, 1},数字次数出现最多的数字是1和2，都出现了3次，第一个1到最后一个1，长度为8，
>第一个2到最后一个2，长度为 6，那么组成最短子数组的是 {2，2，3，1，4，2}，长度为 6
```go
func findShortestSubArray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := len(nums) // 计数
	degree := 1        // 数组的度
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
```

* 448_找到所有数组中消失的数字
```go
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

```
* 442_数组中重复的数据
```go
// 多么适合用哈希表的题...
// 可惜不能使用额外空间，复杂度 O(N)，可能要求在数组原地操作、一次遍历解决
// 解法和下方的 448 很像，充分利用 1~n 的已知条件，十分巧妙
func findDuplicates(nums []int) []int {
	var dups []int
	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		n := num - 1
		if nums[n] < 0 { // 指向的位置以为负值，则重复
			dups = append(dups, num)
			continue
		}
		nums[n] = -nums[n]
	}
	return dups
}
```
* 41_缺失的第一个正数
```go
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num - 1] = -abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
```
* 453_最小操作次数使数组元素相等
```go
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

解法2
先找出数组中的最小值
然后让数组中的每一个值减去最小值,代表着这个数要移动的次数
然后把所有次数累加求和即是结果

func minMoves(nums []int) int {
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
```
* 665_非递减数列 改变1个数使他递增
```go
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
```
* 283_移动零,将0移到最后
```go
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

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
```
* 118_杨辉三角
```go
// 杨辉三角
func generate(numRow int) [][]int {
	ans := make([][]int, numRow)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}

func generate1(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		res = append(res, m)
	}
	return res
}
```
### 链表


### 字符串