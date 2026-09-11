func ceilDiv(a, b int) int {
	if a % b == 0 {
		return a / b
	}
	return a / b + 1
}

func minEatingSpeed(piles []int, h int) int {
	m := piles[0]
	for _, p := range piles {
		m = max(m, p)
	}

	upper := m
	lower := 1
	k := upper

	for ;lower <= upper; {
		mid := (lower + upper) / 2
		sum := 0
		for _, p := range piles {
			sum += ceilDiv(p, mid)
		}

		if sum > h {
			// mid is too small, must be bigger
			lower = mid + 1
		} else {
			// mid is ok, record it
			k = mid
			// but try to find a smaller value
			upper = mid - 1
		}
	}

	return k
}
