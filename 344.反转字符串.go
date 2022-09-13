/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */
package leetcode

// @lc code=start
func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// @lc code=end
