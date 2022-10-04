package stringsExe

func removeElement(nums []int, val int) int {
	lens := len(nums)
	i := 0
	j := 0
	for j < lens {
		if nums[j] == val {
			j++

		}
		nums[i] = nums[j]
		j++
		i++
	}
	return i + 1
}
