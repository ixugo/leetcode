/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */
package leetcode

// @lc code=start
func containsDuplicate(nums []int) bool {
	hmap := make(map[int]int8)
	for _, v := range nums {
		_, ok := hmap[v]
		if ok {
			return true
		}
	}
	return false
}

// @lc code=end
