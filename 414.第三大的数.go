/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */
package leetcode

import (
	"math"
)

// @lc code=start
func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt, math.MinInt, math.MinInt
	for _, v := range nums {
		if v > m1 {
			m1, m2, m3 = v, m1, m2
		} else if v > m2 && v < m1 {
			m2, m3 = v, m2
		} else if v > m3 && v < m2 {
			m3 = v
		}
	}

	if m3 == math.MinInt {
		return m1
	}
	return m3
}

// @lc code=end
