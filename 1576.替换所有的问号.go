/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] 替换所有的问号
 */
package leetcode

// 1. 找出问号
// 2. 取一个不同于问号前后参数的值，替换上去

// @lc code=start
func modifyString(s string) string {
	r := []rune(s)
	for i, v := range r {
		if v != '?' {
			continue
		}

		for x := 'a'; x <= 'c'; x++ {
			ok1 := i > 0 && r[i-1] == x
			ok2 := i < len(r)-1 && r[i+1] == x
			if ok1 || ok2 {
				continue
			}
			r[i] = x
			break
		}

	}

	return string(r)
}

// @lc code=end
