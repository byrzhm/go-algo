package lc_2234

import "testing"

func TestLC2234(t *testing.T) {
	flowers := []int{1, 3, 1, 1}
	newFlowers := int64(7)
	target := 6
	full := 12
	partial := 1
	res := maximumBeauty(flowers, newFlowers, target, full, partial)
	if res != 14 {
		t.Errorf("maximum beauty of flowers is %d, want 14", res)
	}
}
