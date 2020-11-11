/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	var dp = map[int]int{}
	dp[1] = 1 
	dp[2] = 1
	for i := 3; i <= N; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[N]
}
// @lc code=end

