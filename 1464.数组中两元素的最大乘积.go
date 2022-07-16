/*
 * @lc app=leetcode.cn id=1464 lang=golang
 *
 * [1464] 数组中两元素的最大乘积
 */
package leetcode

// 选择排序，仅排出最大的两个值

// @lc code=start
func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if b > a {
		a, b = b, a
	}

	for i := 2; i < len(nums); i++ {
		v := nums[i]
		if v > a {
			a, b = v, a
		} else if v > b {
			b = v
		}

	}
	return (a - 1) * (b - 1)
}

// @lc code=end
