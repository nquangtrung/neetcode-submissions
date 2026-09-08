func twoSum(numbers []int, target int) []int {
	for ptr1, ptr2 := 0, len(numbers) - 1; ptr1 < ptr2; {
		num1 := numbers[ptr1]
		num2 := numbers[ptr2]
		if num1 + num2 == target {
			return []int{ptr1 + 1, ptr2 + 1}
		}

		if num1 + num2 > target {
			ptr2 -= 1
			continue
		}

		if num1 + num2 < target {
			ptr1 += 1
			continue
		}
	}

	return []int{}
}
