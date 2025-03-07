package lc_2597

import "testing"

func TestLC2597(t *testing.T) {
	nums := []int{2, 4, 6}
	k := 2
	res := beautifulSubsets(nums, k)
	if res != 4 {
		t.Errorf("Expected 4, got %d", res)
	}
}
