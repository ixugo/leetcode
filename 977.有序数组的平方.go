/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package leetcode

// 双指针

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	r := make([]int, n)
	i, j := 0, n-1
	k := j
	for i <= j {
		a, b := nums[i]*nums[i], nums[j]*nums[j]
		if a > b {
			r[k] = a
			i++
		} else {
			r[k] = b
			j--
		}
		k--
	}
	return r
}

// @lc code=end
