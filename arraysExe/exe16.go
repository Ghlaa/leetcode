package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	lens := len(nums)
	sort.Ints(nums) //先进行排序
	res := nums[0] + nums[1] + nums[2]
	minValue := math.Abs(float64((nums[0] + nums[1] + nums[2] - target)))

	for i := 0; i < lens-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, lens-1
		for j < k {
			if math.Abs(float64((nums[i] + nums[j] + nums[k] - target))) < minValue {
				minValue = math.Abs(float64(nums[i] + nums[j] + nums[k] - target))
				res = nums[i] + nums[j] + nums[k]
			}
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k != lens-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[i]+nums[j]+nums[k] > target {
				k--
			} else if nums[i]+nums[j]+nums[k] < target {
				j++
			} else {
				return nums[i] + nums[j] + nums[k]
			}
		}
	}
	return res
}
