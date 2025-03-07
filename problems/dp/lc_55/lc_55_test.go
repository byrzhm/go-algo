package lc_55

import "testing"

func TestLC55(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := canJump(nums)
	if res != true {
		t.Error("Expected true, got false")
	}
}
