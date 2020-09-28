package main

import (
	"fmt"
)

type Programmer interface{
	WriteHelloWorld() string
}
type GoProgrammer struct{
}

func (g *GoProgrammer) WriteHelloWorld() string{
	return "fmt.Println(\"Hello World\")"
}
func TestClient(){
	var p Programmer = &GoProgrammer{}
	// var p Programmer	//定义一个接口变量，但实例是GoProgrammer
	// p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
	
}

func main()  {
	TestClient()
}