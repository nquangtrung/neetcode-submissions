func topKFrequent(nums []int, k int) []int {
	bucket := map[int]int{}

	for _, num := range nums {
		if count, ok := bucket[num]; ok {
			bucket[num] = count + 1
			continue
		}

		bucket[num] = 1
	}

	sortedBucket := map[int][]int{}
	for num, count := range bucket {
		if ns, ok := sortedBucket[count]; ok {
			sortedBucket[count] = append(ns, num)
			continue
		}

		sortedBucket[count] = []int{num}
	}


	result := make([]int, k)
	for ptr, currentCount := 0, len(nums); ptr < k && currentCount > 0; currentCount -= 1 {
		if len(sortedBucket[currentCount]) == 0 {
			continue
		}

		for _, n := range sortedBucket[currentCount] {
			result[ptr] = n
			ptr += 1
		}
	}

	return result[:k]
}
