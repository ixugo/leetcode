/*
 * @lc app=leetcode.cn id=1716 lang=golang
 *
 * [1716] 计算力扣银行的钱
 */
package leetcode

// @lc code=start
func totalMoney(n int) int {
	week, day := 1, 0
	total := 0
	for i := 1; i <= n; i++ {
		total += week + day
		day++
		if day == 7 {
			week++
			day = 0
		}
	}
	return total
}

// @lc code=end
