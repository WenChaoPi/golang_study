package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

func TestString()  {
	var s string	//golang的string不是引用和指针类型，而是数值型。默认零值为""
	fmt.Println(s)

	s = "hello"
	fmt.Println(s)
	fmt.Println(len(s))

	// s[1] = '3'	//string是不可变的byte slice，所以不可以像数组一样进行重新赋值
	s = "\xE4\xB8\xA5"	//可以存储任何二进制数据
	fmt.Println(s)
	fmt.Println(len(s))
}
func TestStringToRune()  {
	s := "中华人民共和国"
	for _, c := range s{
		fmt.Printf("%[1]c %[1]x", c)
	}
}

func TestStringFunc(){
	s := "a,b,c"
	parts := strings.Split(s, ",")	//字符串分割
	for _, part := range parts{
		fmt.Printf("%v ", part)
	}

	//字符串连接
	joinS := strings.Join(parts, "-")
	fmt.Println(joinS)
}

func TestStringConv()  {	//字符串转换
	s := strconv.Itoa(10)
	fmt.Println(s, reflect.TypeOf(s))	//使用reflect获取任意变量的类型
	if a, err := strconv.Atoi(s); err == nil{
		fmt.Println(a, reflect.TypeOf(a))
	} else {
		fmt.Println("转换失败")
	}
	
}


func main()  {
	TestString()
	TestStringToRune()
	TestStringFunc()
	TestStringConv()
}