/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */
package leetcode

import "math"

// @lc code=start
func dominantIndex(nums []int) int {
	m1, m2, idx := math.MinInt, math.MinInt, 0
	for i, v := range nums {
		if v > m1 {
			m1, m2 = v, m1
			idx = i
		} else if v > m2 {
			m2 = v
		}
	}

	if m1 < m2*2 {
		return -1
	}
	return idx
}

// @lc code=end
