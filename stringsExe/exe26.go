package stringsExe

func removeDuplicates(nums []int) int {
	hashmap := map[int]int{}
	left := 0
	for i, x := range nums {
		if m, ok := hashmap[x]; !ok { //如果元素x不在hash表里面
			hashmap[x] = i
			nums[left] = m
			left++
		}
	}
	return left
}
