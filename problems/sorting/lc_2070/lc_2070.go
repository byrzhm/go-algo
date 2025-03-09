package lc_2070

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	ilen := len(items)
	prepares := make([]int, ilen+1)
	for i, item := range items {
		prepares[i+1] = max(prepares[i], item[1])
	}

	qlen := len(queries)
	ans := make([]int, qlen)
	for i, query := range queries {
		idx := sort.Search(ilen, func(i int) bool {
			return items[i][0] > query
		})
		ans[i] = prepares[idx]
	}
	return ans
}
