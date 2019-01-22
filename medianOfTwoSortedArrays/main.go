package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {

		if i == len(nums1) {
			nums = append(nums, nums2[j:]...)
			break
		}

		if j == len(nums2) {
			nums = append(nums, nums1[i:]...)
			break
		}

		n1 := nums1[i]
		n2 := nums2[j]
		if n1 <= n2 {
			nums = append(nums, n1)
			i++
			continue
		}

		if n2 < n1 {
			nums = append(nums, n2)
			j++
			continue
		}
	}

	m := math.Floor(float64(len(nums) / 2))
	result := 0.0
	if len(nums)%2 != 0 {
		result = float64(nums[int(m)])
	} else {
		result = float64(nums[int(m)]+nums[int(m)-1]) / 2
	}

	return result
}
