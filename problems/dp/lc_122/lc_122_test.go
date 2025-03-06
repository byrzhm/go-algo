package lc_122

import "testing"

func TestLC122(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	if res != 7 {
		t.Errorf("maxProfit(%d) = %d, want 7", prices, res)
	}
}
