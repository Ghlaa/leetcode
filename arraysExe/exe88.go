package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for j := 0; j < n; j++ {
		nums1[m+j] = nums2[j]
	}
	sort.Ints(nums1)
}
