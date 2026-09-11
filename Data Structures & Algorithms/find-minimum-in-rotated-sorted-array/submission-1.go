func findPivotIndex(nums []int) int {
	left := 0
	right := len(nums) - 1

	for ;left <= right; {
		if left == right {
			return left
		}

		mid := (left + right) / 2

		// fmt.Printf("check %d, %d, %d (%d, %d, %d)\n", left, mid, right, nums[left], nums[mid], nums[right])

		if nums[left] > nums[mid] {
			// pivot point is left of mid
			right = mid
		} else if nums[left] < nums[mid] {
			// pivot point is right of mid
			left = mid
		} else {
			// equal is only when left == mid
			// so right == left + 1
			if nums[right] > nums[left] {
				// pivot point is right
				return right
			} else {
				// pivot point is left
				return left
			}
		}
	}

	return left
}
func findMin(nums []int) int {
	pivot := findPivotIndex(nums)
	
	return nums[(pivot + 1) % len(nums)]
}
