package lc_2070

import "testing"

func TestLC2070(t *testing.T) {
	items := [][]int{
		{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5},
	}
	queries := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 5, 5, 6, 6}

	res := maximumBeauty(items, queries)
	if len(res) != len(expected) {
		t.Errorf("length of results is not equal to expected length")
	}

	for i, ans := range res {
		if ans != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], ans)
		}
	}
}
