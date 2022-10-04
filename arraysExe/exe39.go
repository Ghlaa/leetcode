package main

import "fmt"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	var dfs func(index int, cursum int, tempres []int)
	dfs = func(index int, cursum int, tempres []int) {
		fmt.Println(index, cursum, tempres)
		//胡同
		if cursum == target {
			newTemp := make([]int, len(tempres))
			copy(newTemp, tempres)
			res = append(res, newTemp)
			return
		}
		//剪纸
		if index == n || cursum > target {
			return
		}
		tempres = append(tempres, candidates[index])
		dfs(index, cursum+candidates[index], tempres)
		tempres = tempres[0 : len(tempres)-1]
		dfs(index+1, cursum, tempres)
	}
	dfs(0, 0, []int{})
	return res

}
