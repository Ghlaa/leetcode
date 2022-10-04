package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	lens := len(nums)
	if lens <= 2 {
		return res
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[lens-1] < 0 { //全部为0的话是没有答案的
		return res
	}
	for i := 0; i < lens-1; i++ {
		if nums[i] > 0 { //如果第一个元素都大于0的话，后面肯定不可能相加满足等于0
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, lens-1
		for j < k {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k != lens-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				temp := []int{nums[i], nums[j], nums[k]}
				res = append(res, temp)
				j++
				k--
			}
		}
	}
	return res
}
