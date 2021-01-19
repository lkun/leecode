package main

import "fmt"

func main() {
	len := lengthOfLongestSubstring("abcabcbb")
	fmt.Println("len=", len)
	len1 := lengthOfLongestSubstring1("abcabcbb")
	fmt.Println("len=", len1)
}

func lengthOfLongestSubstring(s string) int {
	maxStrLen := 0
	charMap := make(map[int32]int, 128) //减少内存分配
	p := 0                              //重复字符的上次出现的索引,上一个重复字符和和当前重复字符的之间的子串即有效子串
	for i, char := range s {
		j, ok := charMap[char]
		if ok {
			//字符重复,从重复字符下一位开始重新算子串
			if j > p {
				p = j
			}
		}
		//防止字符串长度为1时,i-p=0
		if i+1-p > maxStrLen {
			maxStrLen = i + 1 - p
		}
		charMap[char] = i + 1
	}
	return maxStrLen
}

func lengthOfLongestSubstring1(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
