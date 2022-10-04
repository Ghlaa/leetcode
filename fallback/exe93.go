package fallback

//
//	import "strconv"
//
//	var res []string
//	func restoreIpAddresses(s string) []string {
//		res=[]string{}
//		dfss(s,0,0,"")
//		return res
//	}
//
//	func dfss(s string,index int,pointNum int,strTemp string){
//		if pointNum==3&&index==len(s){
//			res=append(res,strTemp[:len(strTemp)-1])
//			return
//		}
//		for i:=index;i<len(s);i++{
//			if ishefa(s[index:i+1]){
//				strTemp=strTemp+s[index:i+1]+"."
//				pointNum+=1
//				dfss(s,i+1,pointNum,strTemp)
//				pointNum-=1
//				strTemp=strTemp[0:len(strTemp)-(i+2-index)]
//			}
//
//		}
//	}
//	func ishefa(s string)bool{
//		lens:=len(s)
//		if lens==1{
//			return true
//		}
//		if s[0]=='0'{
//			return false
//		}
//		snit,_:=strconv.Atoi(s)
//		if snit>=0&&snit<=255{
//			return true
//		}else{
//			return false
//		}
//
//	}
