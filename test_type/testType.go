package main
import "fmt"

type MyInt int64
func TestImplicit() {
	var a int = 1
	var b int64
	// b = a	//不支持隐式类型转化
	b = int64(a)

	fmt.Println(a, b)

	var c MyInt
	// c = b
	c = MyInt(b)

	fmt.Println(c, b)

}

func TestPoint(){
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1	//不支持指针运算

	fmt.Println(a, aPtr)
}
func TestString(){
	var s string	//string 是值类型，默认为“”空串， 不是nil
	fmt.Println(s+"is empty string")

}

func main(){
	// TestImplicit()
	// TestPoint()
	TestString()
}

