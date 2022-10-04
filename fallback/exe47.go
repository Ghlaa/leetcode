package fallback

var res [][]int
var hash map[int]bool

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	hash = make(map[int]bool, 0)
	dfs(nums, 0, []int{})
	return res
}
func dfs(nums []int, count int, temp []int) {
	if count == len(nums) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		return
	}
	hash2 := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {

		if hash[i] == false && hash2[nums[i]] == false {
			hash2[nums[i]] = true
			hash[i] = true
			temp = append(temp, nums[i])
			count++
			dfs(nums, count, temp)
			count--
			temp = temp[0 : len(temp)-1]
			hash[i] = false
		}

	}
}
