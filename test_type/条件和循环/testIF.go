package main
import "fmt"

func TestIfContain()  {
	if a := 1==1; a {	//条件中第部分可以给一个声明
		fmt.Println("a is",a)
	}
}

// func main()  {
// 	TestIfContain()
// }