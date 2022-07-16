/*
 * @lc app=leetcode.cn id=1417 lang=golang
 *
 * [1417] 重新格式化字符串
 */
// @lc code=start
package leetcode

func reformat(s string) string {
	l := len(s)
	a, b := make([]rune, 0, l/2+1), make([]rune, 0, l/2+1)
	for _, v := range s {
		if v >= '0' && v <= '9' {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}

	if len(b) > len(a) {
		a, b = b, a
	}
	if len(a)-len(b) > 1 {
		return ""
	}

	r := make([]rune, 0, l)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			r = append(r, a[i/2])
		} else {
			r = append(r, b[i/2])
		}
	}

	return string(r)
}

// @lc code=end
