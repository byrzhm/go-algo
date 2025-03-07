package lc_2597

func beautifulSubsets(nums []int, k int) int {
	ans := -1
	n := len(nums)
	cnt := map[int]int{}

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}

		dfs(i + 1)
		x := nums[i]
		if cnt[x+k] == 0 && cnt[x-k] == 0 {
			cnt[x]++
			dfs(i + 1)
			cnt[x]--
		}
	}
	dfs(0)
	return ans
}
