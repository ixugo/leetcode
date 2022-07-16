/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
package leetcode

// @lc code=start
func lengthOfLastWord(s string) int {

	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}

// @lc code=end
