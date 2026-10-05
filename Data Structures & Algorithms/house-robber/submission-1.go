func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cache := map[int]int{}

	var robHouse func(i int) int
	robHouse = func(i int) int {
		if i >= len(nums) {
			return 0
		}

		if _, found := cache[i]; found {
			return cache[i]
		}

		maxValue := 0
		for j := i + 2; j < len(nums); j++ {
			maxValue = max(maxValue, robHouse(j))
		}

		result := nums[i] + maxValue
		cache[i] = result

		return result
	}

	return max(robHouse(0), robHouse(1))
}
