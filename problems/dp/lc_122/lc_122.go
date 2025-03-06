package lc_122

// Dynamic Programming
/*
func maxProfit(prices []int) int {
	n := len(prices)
	f := make([][2]int, n+1)
	f[0][1] = math.MinInt
	for i := 1; i < n+1; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]+prices[i-1])
		f[i][1] = max(f[i-1][1], f[i-1][0]-prices[i-1])
	}
	return f[n][0]
}
*/

func maxProfit(prices []int) int {
	ans := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
