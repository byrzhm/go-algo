package lc_2234

import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	sort.Ints(flowers)
	idx := sort.SearchInts(flowers, target)
	if idx == 0 {
		return int64(full) * int64(n)
	}
	var leftFlowers int64 = newFlowers - int64(target)*int64(idx)
	for i := 0; i < idx; i++ {
		leftFlowers += int64(flowers[i])
	}

	if leftFlowers >= 0 {
		return max(int64(full)*int64(n), int64(full)*int64(n-1)+int64(partial)*int64(target-1))
	}

	var ans, preSum int64
	j := 0
	for i := 1; i <= idx; i++ {
		leftFlowers += int64(target - flowers[i-1])
		if leftFlowers < 0 {
			continue
		}

		for j < i && int64(flowers[j])*int64(j) <= preSum+leftFlowers {
			preSum += int64(flowers[j])
			j++
		}

		var avg int64 = (leftFlowers + preSum) / int64(j)
		var totalBeauty int64 = avg*int64(partial) + int64(n-i)*int64(full)
		ans = max(ans, totalBeauty)
	}

	return ans
}
