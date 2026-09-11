
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

func search(nums []int, target int) int {
	pivot := findPivotIndex(nums)
	if (nums[pivot] == target) {
		return pivot
	}
	if (nums[0]) == target {
		return 0
	}
	if (nums[len(nums) - 1] == target) {
		return len(nums) - 1
	}
	if (nums[(pivot + 1) % len(nums)] == target) {
		return (pivot + 1) % len(nums)
	}

	left := 0
	right := 0
	if nums[0] < target && target < nums[pivot] {
		left = 0
		right = pivot
	} else {
		left = pivot + 1 % len(nums)
		right = len(nums) - 1
	}

	for ;left <= right; {
		fmt.Printf("searching between %d, %d (%d, %d)\n", left, right, nums[left], nums[right])
		if left == right {
			if nums[left] != target {
				return -1
			} else {
				return left
			}
		}

		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
