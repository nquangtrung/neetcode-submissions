func permute(nums []int) [][]int {
	picked := map[int]bool{}
	result := [][]int{}
	current := []int{}

	var pick func(i int)
	pick = func(i int) {
		current = append(current, nums[i])
		picked[i] = true
		if len(current) == len(nums) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		for j, _ := range nums {
			if _, ok := picked[j]; ok {
				continue
			}

			pick(j)
			current = current[:len(current) - 1]
			delete(picked, j)
		}
	}

	for i, _ := range nums {
		pick(i)
		current = []int{}
		picked = map[int]bool{}
	}

	return result
}
