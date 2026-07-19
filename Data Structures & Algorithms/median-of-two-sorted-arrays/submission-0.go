func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	totalLen := (len(nums1) + len(nums2))

	medianPos := (totalLen + 1) / 2
	l := 0
	r := len(nums1)
	for l <= r {
		mid1 := l + (r-l)/2
		mid2 := medianPos - mid1

		l1 := math.MinInt32
		l2 := math.MinInt32
		r1 := math.MaxInt32
		r2 := math.MaxInt32

		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}
		if mid1 < len(nums1) {
			r1 = nums1[mid1]
		}

		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}
		if mid2 < len(nums2) {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			if totalLen%2 == 0 {
				return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2.0
			} else {
				return float64(max(float32(l1), float32(l2)))
			}
		} else if l1 > r2 {
			r = mid1 - 1
		} else {
			l = mid1 + 1
		}
	}
	return 0
}
