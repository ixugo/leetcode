/*
 * @lc app=leetcode.cn id=1991 lang=golang
 *
 * [1991] 找到数组的中间位置
 */
package array

// @lc code=start
func findMiddleIndex(nums []int) int {
	left, right := 0, 0
	for _, v := range nums {
		right += v
	}

	for i, v := range nums {
		right -= v
		if right == left {
			return i
		}
		left += v
	}
	return -1
}

// @lc code=end
