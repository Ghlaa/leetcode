package main

import (
	"fmt"

	"sort"
)

func nextPermutation(nums []int) {
	flag := -1 //记录下标变量
	//遍历数组
	for i := len(nums) - 1; i > 0; i-- {
		//找到第一个满足右边大于左边的数字下标
		if nums[i] > nums[i-1] {
			//如果找到就记录下标
			flag = i
			//寻找记录下标往右序列的大于target的最小值的索引
			index := targetIndex(nums[flag:], nums[flag-1])
			//进行位置交换
			temp := nums[i+index]
			nums[i+index] = nums[i-1]
			nums[i-1] = temp
			break
		}
	}
	//如果flag没有赋值，说明此时排列已经是最大得了
	if flag == -1 {
		sort.Ints(nums)
	} else {
		//否则对flag往右的序列进行升序排序就行
		sort.Ints(nums[flag:])
	}
	fmt.Println(flag, nums)
}

//返回大于target中最小的数的索引
func targetIndex(shuzu []int, target int) int {
	resindex := 0
	minBalcance := 101 //因为数组中的范围为0到100，所以任意两个数的差小于101
	for i := 0; i < len(shuzu); i++ {
		//前提是这个数字要大于target
		if shuzu[i] > target {
			//寻找大于target最小的数，返回其下标索引
			if abs(shuzu[i]-target) <= minBalcance {
				minBalcance = abs(shuzu[i] - target)
				resindex = i
			}
		}
	}
	return resindex
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}
