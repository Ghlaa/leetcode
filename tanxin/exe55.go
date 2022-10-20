package tanxin
func jump(nums []int) int {
	lens:=len(nums)
	max:=0
	step:=1
	for i:=0;i<=max;i++{

		if i+nums[i]>max{
			max=i+nums[i]
			step++
		}
		if max>=lens-1{
			return step
		}
	}
	return step
}