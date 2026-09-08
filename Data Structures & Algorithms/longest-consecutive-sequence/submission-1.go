func longestConsecutive(nums []int) int {
	posMap := make(map[int]int)

	for idx, num := range nums {
		posMap[num] = idx
	}

	maxCount := 0
	for _, num := range nums {
		if _, isNotStartOfSeq := posMap[num - 1]; isNotStartOfSeq {
			// num is not start of seq			
			continue
		}

		count := 0
		current := num
		for {
			_, ok := posMap[current]
			if !ok {
				break;
			}
			count += 1
			current += 1
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
