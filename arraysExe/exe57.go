package main

func insert(intervals [][]int, newInterval []int) [][]int {
	//结果集
	res := [][]int{}
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	i := 0
	//遍历区间
	for ; i < len(intervals); i++ {
		//如果当前区间的结束位置小于插入区间的开始位置,那么可以直接加入结果集
		if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		} else {
			break
		}
	}
	//遍历集合,对插入区间进行合并
	for ; i < len(intervals); i++ {
		//如果当前区间的起始位置小于插入区间的结束位置，考虑合并
		if intervals[i][0] <= newInterval[1] {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		} else {
			//否则退出，将剩余的区间直接添加到结果集
			break
		}
	}
	res = append(res, newInterval)  //把合并后的区间加入到结果集
	for ; i < len(intervals); i++ { //剩余区间都是大于插入区间的，可以直接加入到结果集
		res = append(res, intervals[i])
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
