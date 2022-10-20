package main

//var  mapstr map[int]string
//var  ress []string
//func letterCombinations(digits string) []string {
//	mapstr=make(map[int]string,10)
//	mapstr[0]=""
//	mapstr[1]=""
//	mapstr[2]="abc"
//	mapstr[3]="def"
//	mapstr[4]="ghi"
//	mapstr[5]="jkl"
//	mapstr[6]="mno"
//	mapstr[7]="pqrs"
//	mapstr[8]="tuv"
//	mapstr[9]="wxyz"
//	ress=[]string{}
//	lens:=len(digits)
//	if lens==0{
//		return ress
//	}
//	fallack3(digits,0,lens,"")
//	return ress
//}
//
//func fallack3(digits string,start int,lens int,str string){
//	if start==lens{
//		ress=append(ress,str)
//		return
//	}
//	strtemp:=string(digits[start])
//	digitsValue,_:=strconv.Atoi(strtemp)
//
//	mapVaueStr:=mapstr[digitsValue]
//	for i:=0;i<len(mapVaueStr);i++{
//		str+=string(mapVaueStr[i])
//		fallack3(digits,start+1,lens,str)
//		str=str[0:len(str)-1]
//	}
//}