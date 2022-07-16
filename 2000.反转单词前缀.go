/*
 * @lc app=leetcode.cn id=2000 lang=golang
 *
 * [2000] 反转单词前缀
 */
package leetcode

// @lc code=start
func reversePrefix(word string, ch byte) string {
	r := []rune(word)
	idx := -1
	for i, v := range word {
		if v == rune(ch) {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}

	i := 0
	for i <= idx {
		r[i], r[idx] = r[idx], r[i]
		i++
		idx--
	}

	return string(r)
}

// @lc code=end
