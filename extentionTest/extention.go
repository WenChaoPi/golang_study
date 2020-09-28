package main

import (
	"fmt"
)
//定义 Pet
type Pet struct{
}
func (p *Pet) Speak(){
	fmt.Println("...")
}
func (p *Pet) SpeakTo(host string){
	p.Speak()
	fmt.Println(" ", host)
}

//定义 Dog
// type Dog struct{
// 	p *Pet
// }
// func (d *Dog) Speak(){
// 	fmt.Println("Wang!")
// }
// func (d *Dog) SpeakTo(host string){
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

// func TestDog(){
// 	dog := new(Dog)
// 	dog.SpeakTo("Chao")
// }

//支持隐式组装
type Dog struct{
	Pet
}
func TestDog(){
	dog := new(Dog)
	dog.SpeakTo("chao")
}
func (d *Dog) Speak(){
	fmt.Println("Wang!")
}

/*注意：
1、内嵌但结构类型不能当作继承来用
2、虽然上述Dog重写了Speak方法，但TestDog（）仍然输出 ...
*/



type Code string	//自定义类型
type Programmer interface{	//定义Programmer接口
	WriteHelloWorld() Code
}
//
type GoProgrammer struct{	//GoProgrammer接口实例
}
func (p *GoProgrammer) WriteHelloWorld() Code{
	return "fmt.Println(\"Hello Golang\")"
}

type JavaProgrammer struct{	//JavaProgrammer接口实例
}
func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "fmt.Println(\"Hello Java\")"
}

//
func WriteFirstProgram(p Programmer){	//注意这里参数是接口interface， 对应下面的调用一定是指针类型
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(){
	// goPro := GoProgrammer{}	//这种写法错误，应该写成 goPro := &GoProgrammer{}
	goPro := new(GoProgrammer)
	javaPro := new(JavaProgrammer)

	WriteFirstProgram(goPro)
	WriteFirstProgram(javaPro)
}



func main()  {
	// TestDog()
	TestPolymorphism()
}