/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	data := make(map[int]int, len(nums))
	for i, v := range nums {
		j, ok := data[target-v]
		if ok {
			return []int{i, j}
		}
		data[v] = i
	}
	return nil
}

// @lc code=ends
