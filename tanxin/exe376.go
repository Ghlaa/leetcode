package tanxin
func wiggleMaxLength(nums []int) int {
	flag:=0
	lens:=len(nums)
	if lens==0||lens==1{
		return lens
	}
	temp:=nums[0]
	count:=1
	if nums[1]>nums[0]{
		flag=1
		temp=nums[1]
		count++
	}else if nums[1]-nums[0]<0{
		flag=-1
		temp=nums[1]
		count++
	}

	for i:=2;i<len(nums);{
		if flag==1&&nums[i]<temp{
			temp=nums[i]
			i++
		}else{
			flag=-1
			count++
		}
		if flag==-11&&nums[i]>temp{
			temp=nums[i]
			i++
		}else{
			flag=1
			count++
		}
	}
	return count
}