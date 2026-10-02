func climbStairs(n int) int {
	cache := map[int]int{}

	var climb func (n int) int
	climb = func(n int) int {
		if n <= 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}

		if result, found := cache[n]; found {
			return result
		}

		result := climb(n - 1) + climb(n - 2)
		cache[n] = result

		return result
	}

	return climb(n)
}
