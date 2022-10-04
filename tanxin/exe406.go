package tanxin

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
	})
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1]+1:], result[info[1]:]) //将插入位置之后的元素后移动一位（意思是腾出空间）
		result[info[1]] = info                     //将插入元素位置插入元素
	}
	return result
}
