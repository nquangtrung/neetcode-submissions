func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array to minimize binary search range
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partition1 := (low + high) / 2
		partition2 := (m+n+1)/2 - partition1

		// If partition1 is 0, nothing is on the left side of nums1
		maxLeft1 := math.MinInt64
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		// If partition1 is m, nothing is on the right side of nums1
		minRight1 := math.MaxInt64
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		// If partition2 is 0, nothing is on the left side of nums2
		maxLeft2 := math.MinInt64
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		// If partition2 is n, nothing is on the right side of nums2
		minRight2 := math.MaxInt64
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		// Check if we found the correct partition
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			// Odd total length
			if (m+n)%2 != 0 {
				return float64(max(maxLeft1, maxLeft2))
			}
			// Even total length
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
		} else if maxLeft1 > minRight2 {
			// We are too far right in nums1, move left
			high = partition1 - 1
		} else {
			// We are too far left in nums1, move right
			low = partition1 + 1
		}
	}

	return 0.0
}