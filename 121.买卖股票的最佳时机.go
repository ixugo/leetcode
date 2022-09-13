/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package leetcode

// 贪心算法
// 找到最左最小值，将后面的值-最小值判断最大收益

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxIncome, minV := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		v := prices[i]
		if v <= minV {
			minV = v
			continue
		}
		if t := v - minV; t > maxIncome {
			maxIncome = t
		}
	}
	return maxIncome
}

// @lc code=end
