package main

import "fmt"

func main() {
	fmt.Println(twoSum1([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum3([]int{3, 3}, 6))         // [0,1]
}

//
// 暴力双重遍历
//
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ { // 注意j的遍历区间
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//
// 哈希表索引
// 检查数组中是否存在目标元素，若存在则找出索引
// 哈希表特别适合抽象类的配对结构，当要解决问题的数据单元是成对数据关系时，考虑哈希表 map 结构
//
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	// 先全部将数组中的每个元素塞到map中，key：数组的值，value：数组下标
	for i, num := range nums {
		m[num] = i
	}
	// 循坏，目标(Target)-当前循环值(A)=需要在map中找是否存在的值(B)，如果存在，那么A+B=Target
	for i, num := range nums {
		pair := target - num
		j, ok := m[pair]
		if ok && i != j { // 剔除自身相加的情况，使用哈希表需要注意的点：索引重叠时为同一元素
			return []int{i, j}
		}
	}
	return nil
}

//
// 优化后的哈希表索引方式
// 看 twoSum2 会发现进行了2次for循环，可以进行合并优化，一边遍历，一边检查
//
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
