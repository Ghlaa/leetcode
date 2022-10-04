package stringsExe

//两数之和
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}

func twoSum1(nums []int, target int) []int {
	for i, m := range nums {
		for j := i + 1; j < len(nums); j++ {
			if m+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, m := range nums {
		if j, ok := hashmap[target-m]; ok {
			return []int{j, i}
		}
		hashmap[m] = i
	}
	return nil
}
