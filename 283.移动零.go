/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package leetcode

// @lc code=start
func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}

// @lc code=end
