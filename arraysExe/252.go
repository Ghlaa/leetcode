package main

import (
	"sort"
)

func canAttendMeetings(temp [][]int) bool {
	//将区间按照会议的开始顺序进行排序
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][0] < temp[j][0]
	})
	falg := true
	//遍历会议，如果下一个会议在前一个会议之前就结束了，返回false
	for i := 0; i < len(temp); i++ {
		if i > 0 && temp[i][0] < temp[i-1][1] {
			return false
		}
	}
	return falg
}
