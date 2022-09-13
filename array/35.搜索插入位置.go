/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package array

// @lc code=start
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {

		mid := (j-i)/2 + i
		v := nums[mid]
		if v == target {
			return mid
		}

		if v < target {
			i = mid + 1
		} else {
			j = mid - 1
		}

	}
	return i
}

// @lc code=end
