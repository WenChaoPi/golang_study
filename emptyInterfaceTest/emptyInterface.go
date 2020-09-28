package main

import(
	"fmt"
)

func DoSomething(p interface{})  {
	if i, ok := p.(int); ok{	//类型断言
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok{
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}
func SwitchDoSomething(p interface{})  {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}
type MyInt int64
func TestEmptyInterfaceAssertion(){
	DoSomething(10)
	DoSomething("abc")
	var a MyInt = 12
	DoSomething(a)
}
func TestEmptyInterface(){
	SwitchDoSomething(10)
	SwitchDoSomething("10")
	var a MyInt = 12
	SwitchDoSomething(a)
}

func main()  {
	// TestEmptyInterfaceAssertion()	
	TestEmptyInterface()
}