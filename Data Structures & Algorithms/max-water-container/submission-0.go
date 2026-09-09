func maxArea(heights []int) int {
	result := 0
	i := 0
	j := len(heights) - 1
	for {
		if i >= j {
			break
		}

		area := (j - i) * min(heights[i], heights[j])
		if (area > result) {
			result = area
		}

		if heights[i] < heights[j] {
			i += 1
		} else if heights[i] > heights[j] {
			j -= 1
		} else {
			i += 1
		}
	}
	return result 
}
