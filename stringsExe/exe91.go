package stringsExe

//func numDecodings(s string) int {
//lens:=len(s)
//if lens==0{
//return 0
//}
//if s[0]-'0'==0{
//return 0
//}
//if lens==1{
//return 1
//}
//
//dp:=[101]int{}
//dp[0]=1
//if s[0]-'0'<=2&&(s[1]-'0')==0{//已经确定第一位不是0了，并且至少是两位
//dp[1]=1
//}else if s[0]-'0'<=2&&(s[1]-'0')<=6{//已经确定第一位不是0了，并且至少是两位
//dp[1]=2
//}else if s[0]-'0'<=2&&(s[1]-'0')>6{//已经确定第一位不是0了，并且至少是两位
//dp[1]=1
//}else if s[0]-'0'>2{
//dp[1]=1
//}
////for i:=lens-1;i>=2;i--{
//// if s[i]-'0'==0{
//// 	if s[i-1]-'0'>2||s[i-1]-'0'==0{
//// 		return 0
//// 	}else {
//// 		dp[i]=dp[i-2]
//// 	}
//// }
//// else if s[i-1]-'0'<=2&&s[i]-'0'<=6{
//// 	dp[i]=dp[i-1]+1
//// }else {
//// 	dp[i]=dp[i-1]
//// }
//if s[i]-'0'==0{
//dp[i]=dp[i-2]
//}else if s[i]-'0'>6{
//dp[i]=dp[i-1]
//}else{
//if s[i-1]-'0'==0{
//dp[i]=dp[i-1]
//}else if s[i-1]-'0'>2{
//dp[i]=dp[i-1]
//}else{
//dp[i]=dp[i-1]+1
//}
//}
//
//return dp[lens-1]
//}
