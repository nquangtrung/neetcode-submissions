func minCostClimbingStairs(cost []int) int {
	cache := map[int]int{}

	var climb func(n int) int
	climb = func(n int) int {
		if n >= len(cost) {
			return math.MaxInt
		}
		if n == len(cost) - 1 {
			return cost[len(cost) - 1]
		}
		if n == len(cost) - 2 {
			return cost[len(cost) - 2]
		}

		if result, found := cache[n]; found {
			return result
		}

		result := cost[n] + min(climb(n + 1), climb(n + 2))
		cache[n] = result

		return result
	}

	return min(climb(0), climb(1))
}
