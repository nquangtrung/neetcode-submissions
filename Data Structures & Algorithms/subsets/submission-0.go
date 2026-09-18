func subsets(nums []int) [][]int {
	result := [][]int{{}}

	current := []int{}

	var pick func(pos int) 
	pick = func(pos int) {
		current = append(current, nums[pos])
		result = append(result, append([]int{}, current...))
		// fmt.Printf("current %v, result %v\n", current, result)
		for i := pos + 1; i < len(nums); i++ {
			pick(i)
			current = current[:len(current) - 1]
		}
	}

	for i := range len(nums) {
		pick(i)
		current = []int{}
	}
	
	return result	
}