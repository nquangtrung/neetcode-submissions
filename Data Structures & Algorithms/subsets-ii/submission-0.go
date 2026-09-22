import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{{}}
	current := []int{}

	var pick func (i int)
	pick = func (i int) {
		current = append(current, nums[i])
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for j := i + 1; j < len(nums); {
			pick(j)
			current = current[:len(current) - 1]
			lastNum := nums[j]
			for j < len(nums) && nums[j] == lastNum {
				j++
			}
		}
	}

	for i := 0; i < len(nums); {
		pick(i)
		current = []int{}

		lastNum := nums[i]
		for i < len(nums) && nums[i] == lastNum {
			i++
		}
	}

	return result
}
