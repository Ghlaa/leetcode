package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	res := [][]int{}
	if len(nums) <= 2 {
		return res
	}
	sort.Ints(nums) //-4 -1 -1 0 1 2
	lens := len(nums) - 1
	for i := 0; i < lens-1; i++ {
		if nums[i] > 0 {
			return res
		}

		target := -1 * nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, lens
		for j < k {
			if nums[j]+nums[k] == target {
				temp := []int{nums[i], nums[j], nums[k]}
				res = append(res, temp)
				j++
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[j]+nums[k] < target {
				i++
			} else {
				j--
			}

		}

	}
	return res
}
