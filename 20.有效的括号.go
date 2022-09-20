/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package leetcode

// @lc code=start
func isValid(s string) bool {
	n := len(s) % 2
	if n != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		v := s[i]
		char, ok := pairs[v]
		if !ok {
			stack = append(stack, v)
			continue
		}

		l := len(stack)
		if l == 0 || stack[l-1] != char {
			return false
		}
		stack = stack[:l-1]
	}
	return len(stack) == 0
}

// @lc code=end
