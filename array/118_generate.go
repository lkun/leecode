package main

import "fmt"

func main() {
	ans := generate1(5)
	fmt.Println(ans)
}

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
