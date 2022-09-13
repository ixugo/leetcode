/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
package array

// @lc code=start

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		val := strs[0][i]
		for _, v := range strs {
			if len(v) <= i || v[i] != val {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

// @lc code=end
