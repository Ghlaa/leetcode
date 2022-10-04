package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	lens := len(nums)
	if lens <= 3 {
		return res
	}
	sort.Ints(nums)
	if target > 0 && nums[0] >= target { //如果第一个大于target，不满足,
		return res
	}
	if target < 0 && nums[lens-1] <= target { //如果第一个大于target，不满足
		return res
	}
	for i := 0; i < lens-2; i++ {
		if target > 0 && nums[i] >= target { //退出条件
			break
		}
		if i > 0 && nums[i] == nums[i-1] { //去重处理
			continue
		}
		for j := i + 1; j < lens-1; j++ {

			if j != i+1 && nums[j] == nums[j-1] { //去重处理
				continue
			}
			k, l := j+1, lens-1
			for k < l {
				if k != j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				if l != lens-1 && nums[l] == nums[l+1] {
					l--
					continue
				}
				if nums[i]+nums[j]+nums[k]+nums[l] < target {
					k++
				} else if nums[i]+nums[j]+nums[k]+nums[l] > target {
					l--
				} else {
					temp := []int{nums[i], nums[j], nums[k], nums[l]}
					res = append(res, temp)
					k++
					l--
				}
			}
		}
	}
	return res
}
