package stringsExe

import "fmt"

func main() {
	//
	//b:=generateParenthesis(1)
	//fmt.Print(b)
	//fmt.Println("ddd")s

	//var c='2'
	//fmt.Println(c)
	//fmt.Printf("%c\n",c)
	//fmt.Println(string(c))
	//fmt.Println(int(c))
	//fmt.Println(strconv.ParseInt(string(c),10,64))
	//
	var c = "123456"
	fmt.Println(c[0])
	fmt.Println(string(c[0]))
	fmt.Println(byte(c[0]))
	fmt.Println(int(c[0]))
	for _, ch := range c {
		fmt.Println(byte(ch))
	}

	//res:=convert("PAYPALISHIRING", 3)
	//fmt.Println(res)go
	//res:=make([]int,4)
	//res1:=[]int{}
	//fmt.Println(res)
	//fmt.Println(res1)

	//fmt.Println(numberToWords(1234567))
}
