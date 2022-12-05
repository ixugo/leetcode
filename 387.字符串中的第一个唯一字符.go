/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
package leetcode

// @lc code=start
func firstUniqChar(s string) int {
	data := [26]int{}

	for _, v := range s {
		data[v-'a']++
	}

	for i, v := range s {
		if data[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
