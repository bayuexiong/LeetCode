/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] > -1 {
				if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}

// @lc code=end

