import "slices"

func combinationSum2(nums []int, target int) [][]int {
	result := [][]int{}
	sum := 0
	current := []int{}

	slices.Sort(nums)

	var pick func(i int)
	pick = func(i int) {
		// fmt.Printf("pick %d current %v\n", nums[i], current)
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

		for j := i + 1; j < len(nums); {
			pick(j)

			current = current[:len(current) - 1]
			sum -= nums[j]

			lastNum := nums[j]
			for j < len(nums) && nums[j] == lastNum {
				j++
			}
		}
	}

	for i := 0; i < len(nums); {
		pick(i)
		current = []int{}
		sum = 0

		lastNum := nums[i]
		for i < len(nums) && nums[i] == lastNum {
			i++
		}
	}

	return result
}
