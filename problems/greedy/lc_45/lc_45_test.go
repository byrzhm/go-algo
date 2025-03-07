package lc_45

import "testing"

func TestLC45(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := jump(nums)
	if res != 2 {
		t.Errorf("jump(%d)=%d, want 2", nums, res)
	}
}
