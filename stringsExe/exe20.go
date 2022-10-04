package stringsExe

//func isValid(s string) bool {
//	if len(s)%2==1{
//		return false
//	}
//	hashmap:=make(map[byte]byte)
//	hashmap['(']=')'
//	hashmap['{']='}'
//	hashmap['[']=']'
//	stack:=[]byte{}
//	for i:=0;i<len(s);i++{
//		if hashmap[s[i]]==0{//如果匹配的是右闭合的
//		if len(stack)>0&&hashmap[stack[len(stack)-1]]==s[i]{
//			stack=stack[:len(stack)-1]
//		}else{
//			return false
//		}
//		}else{
//			stack=append(stack,s[i])
//		}
//	}
//	//fmt.Printf("dd", stack)
//	if len(stack)==0 {
//		return true
//	}else{
//		return false
//	}
//}
