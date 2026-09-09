import (
	"slices"
	// "log"
)

func threeSum(nums []int) [][]int {
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})

	// log.Printf("sorted %v", nums)

	result := [][]int{}
	for i := 0; i < len(nums) - 2; {
		target := 0 - nums[i]

		j := i + 1
		k := len(nums) - 1

		for {
			if j >= k {
				break
			}
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				currentK := nums[k]
				for ;k >= 0 && nums[k] == currentK; k-- { }

				currentJ := nums[j]
				for ;j < len(nums) - 1 && nums[j] == currentJ; j++ { }
			} else if sum > target {
				currentK := nums[k]
				for ;k >= 0 && nums[k] == currentK; k-- { }
			} else {
				currentJ := nums[j]
				for ;j < len(nums) - 1 && nums[j] == currentJ; j++ { }
			}
		}
		currentI := nums[i]
		for ;nums[i] == currentI && i < len(nums) - 2; { i++ }
	}
	return result
}
