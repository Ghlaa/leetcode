package main

import (
	"sort"
)

func merged(intervals [][]int) [][]int {
	//定义结果集
	res := [][]int{}
	//先把区间按照起始位置升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0]) //先把第一个区间直接加入到结果集
	//遍历区间
	for i := 1; i < len(intervals); i++ {
		//如果当前区间的起始位置>结果集最后一个区间的最终位置，说明不重叠，直接把当前区间加入到结果集
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else {
			//反之说明重叠，把当前区间合并到结果集的最后一个区间
			res[len(res)-1][1] = maxx(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

//比较大小
func maxx(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
