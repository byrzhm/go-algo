package lc_45

func jump(nums []int) int {
	n := len(nums)
	ans := 0
	currRight := 0
	nextRight := 0
	for i := 0; i < n-1; i++ {
		nextRight = max(nextRight, nums[i]+i)
		if currRight == i {
			currRight = nextRight
			ans++
		}
	}
	return ans
}
