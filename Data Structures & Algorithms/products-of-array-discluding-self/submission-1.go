func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums) + 1)
	rightProduct := make([]int, len(nums) + 1)

	for idx := 0; idx <= len(nums); idx++ {
		if idx == 0 {
			leftProduct[idx] = 1
			rightProduct[idx] = 1
			continue
		}

		leftProduct[idx] = leftProduct[idx - 1] * nums[idx - 1]
		rightProduct[idx] = rightProduct[idx - 1] * nums[len(nums) - idx]
	}

	result := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		result[idx] = leftProduct[idx] * rightProduct[len(nums) - idx - 1]
	}

	return result
}
