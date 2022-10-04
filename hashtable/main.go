package main

import "fmt"

func main() {

	//在读程序之前读者,可以先思考这四行代码输出什么内容
	fmt.Println(string(97)) //前置，输出是a
	fmt.Println(string(20320))
	temp := []rune{20320, 22909, 32, 19990, 30028}
	fmt.Println(string(temp))

	var str string = "hello world"
	fmt.Println("byte=", []byte(str))
	fmt.Println("byte=", []rune(str))
	fmt.Println(str[:2])
	fmt.Println(string([]rune(str)[:2]))

	var str2 string = "你好 世界"
	fmt.Println("byte=", []byte(str2))
	fmt.Println("byte=", []rune(str2))
	fmt.Println(str2[:2])
	fmt.Println(string([]rune(str2)[:2]))
}
