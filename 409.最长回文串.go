/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) int {
	if l := len(s); l < 1 {
		return l
	}

	data := make(map[rune]int)
	for _, v := range s {
		data[v]++
	}

	var total int

	for _, v := range data {
		total += v / 2 * 2
	}
	if total < len(s) {
		total++
	}

	return total
}

// @lc code=end
