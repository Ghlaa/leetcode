package tanxin

func jump(nums []int) int {
	curDis := 0
	nextDis := 0
	res := 0
	lens := len(nums)
	for i := 0; i <= lens; i++ {
		nextDis = max(i+nums[i], nextDis)
		if i == curDis {
			res++
		}
		if nextDis >= lens-1 {
			return res
		}
	}
	return res
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
