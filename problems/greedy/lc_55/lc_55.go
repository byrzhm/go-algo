package lc_55

func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i, x := range nums {
		if i <= rightmost {
			rightmost = max(rightmost, i+x)
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}
