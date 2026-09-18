import "slices"

func combinationSum(nums []int, target int) [][]int {
	result := [][]int{}
	sum := 0
	current := []int{}

	slices.Sort(nums)

	var pick func(i int)
	pick = func(i int) {
		current = append(current, nums[i])
		sum += nums[i]

		if sum == target {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
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
