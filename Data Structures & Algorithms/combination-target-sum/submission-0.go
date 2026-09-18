func combinationSum(nums []int, target int) [][]int {
	result := [][]int{}
	sum := 0
	current := []int{}

	var pick func(i int)
	pick = func(i int) {
		current = append(current, nums[i])
		sum += nums[i]

		// fmt.Printf("current %v %d\n", current, sum)

		if sum == target {
			result = append(result, append([]int{}, current...))
			return
		} else if sum > target {
			return
		}

		for j := i; j < len(nums); j++ {
			pick(j)

			current = current[:len(current) - 1]
			sum -= nums[j]
		}
	}

	for i := 0; i < len(nums); i++ {
		pick(i)
		current = []int{}
		sum = 0
	}

	return result
}
