package stringsExe

//var resstrs []string
//func generateParenthesis(n int) []string {
//	resstrs=[]string{}
//	//resstr:=""
//
//	if n<=0{
//		return resstrs
//	}
//	getStrings("",n,n)
//	return resstrs
//}
//
//
//func getStrings(s string,left,right int) {
//	if left==0&&right==0{
//		resstrs=append(resstrs,s)
//		return
//	}
//	if left==right{
//		getStrings(s+"(",left-1,right)
//	}else if left<right{
//		if left>0{
//			getStrings(s+"(",left-1,right)
//		}
//		getStrings(s+")",left,right-1)
//	}
//}

//
//var resstrs=[]string{}
//func generateParenthesis(n int) []string {
//	resstrs=[]string{}
//	if n<=0{
//		return resstrs
//	}
//	getstring01(0,n,"")
//	return resstrs
//}
//func getstring01(index,n int,s string){
//	if index==2*n{
//		if isValidstr(s){
//			resstrs=append(resstrs,s)
//		}
//		return
//	}
//	getstring01(index+1,n,s+"(")
//	getstring01(index+1,n,s+")")
//}
//
//func isValidstr(s string)bool {
//	stack:=0
//	for i:=0;i<len(s);i++{
//		if s[i]=='('{
//			stack++
//		}else if s[i]==')'{
//			stack--
//		}
//		if stack<0{
//			return false
//		}
//	}
//	if stack==0{
//		return true
//	}else{
//		return false
//	}
//}

var resstrs = []string{}

func generateParenthesis(n int) []string {
	resstrs = []string{}
	if n <= 0 {
		return resstrs
	}
	getstring01(0, 0, n, "")
	return resstrs
}
func getstring01(index, stack, n int, s string) {

	if index == 2*n && stack == 0 {
		resstrs = append(resstrs, s)
		return
	}
	if stack == 0 {
		getstring01(index+1, stack+1, n, s+"(")
	}
	if stack > 0 {
		getstring01(index+1, stack-1, n, s+")")
		getstring01(index+1, stack+1, n, s+"(")
	}

}
