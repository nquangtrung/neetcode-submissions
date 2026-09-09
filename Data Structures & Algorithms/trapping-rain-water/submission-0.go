func trap(height []int) int {
	maxPrefix := make([]int, len(height))
	maxSuffix := make([]int, len(height))

	max := 0
	for i, h := range height {
		if h > max {
			max = h
		}

		maxPrefix[i] = max
	}

	max = 0
	for i := len(height) - 1; i >= 0; i-- {
		h := height[i]
		if h > max {
			max = h
		}

		maxSuffix[i] = max
	}

	// fmt.Printf("max %v, %v\n", maxPrefix, maxSuffix)

	total := 0

	for i, h := range height {
		level := min(maxPrefix[i], maxSuffix[i]) - h
		total += level
	}

	return total
}
