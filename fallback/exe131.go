package main

/*import "fmt"

var res [][]string
func partition(s string) [][]string {
	res=[][]string{}
	dfs(s,0,[]string{})
	return res
}
func dfs(s string,index int,temp []string){
	//剪纸
	if ishuiwen(temp)==false{
		return
	}
	if index==len(s)&&ishuiwen(temp)==true{
		fmt.Println(temp)
		tmpString:=make([]string,len(temp))
		copy(tmpString,temp)
		res=append(res,tmpString)
		return
	}
	for i:=index;i<len(s);i++{
		temp=append(temp,s[index:i+1])
		fmt.Println(temp)
		dfs(s,i+1,temp)
		temp=temp[0:len(temp)-1]
	}
}

func ishuiwen(temp []string)bool{
	lens:=len(temp)
	for i:=0;i<lens;i++{
		if ish(temp[i])==false{
			return false
		}
	}
	return true
}

func ish(s string)bool{
	lens:=len(s)
	for i:=0;i<lens/2;i++{
		if s[i]!=s[lens-i-1]{
			return false
		}
	}
	return true
}*/