package main

var res [][]string
func solveNQueens(n int) [][]string {
	res=make([][]string,0)
	restmp:=make([]string,0)
	initres(n,&restmp)
	//fmt.Println(restmp)
	dfs(n,restmp,0)
	return res
}

func initres(n int,restmp *[]string){
	tempstr:=""
	for i:=0;i<n;i++{
		tempstr+="."
	}
	//fmt.Println(tempstr)
	for i:=0;i<n;i++{
		*restmp=append(*restmp,tempstr)
	}
	//fmt.Println(*restmp)
}

func dfs(n int,restmp []string,row int){
	if row==n{
		tmp:=make([]string,n)
		copy(tmp,restmp)
		res=append(res,tmp)
		//fmt.Println("restmpr",restmp)
		return
	}
	for i:=0;i<n;i++{
		if isvalid(n,restmp,row,i){
			slicestr:=[]byte(restmp[row])//无法直接修改一个字符串的值
			slicestr[i]='Q'
			restmp[row]=string(slicestr)
			row++
			dfs(n,restmp,row)
			row--
			slicestr=[]byte(restmp[row])//无法直接修改一个字符串的值,
			slicestr[i]='.'
			restmp[row]=string(slicestr)
		}
	}
}

func isvalid(n int,restmp []string,row int,index int)bool{
	for i:=0;i<row;i++{//竖线
		if restmp[i][index]=='Q'{
			return false
		}
	}
	for i,j:=row-1,index-1;i>=0&&j>=0;i,j=i-1,j-1{//135
		if restmp[i][j]=='Q'{
			return false
		}
	}
	for i,j:=row-1,index+1;i>=0&&j<n;i,j=i-1,j+1{
		if restmp[i][j]=='Q'{
			return false
		}
	}
	return true
}