package main
import "fmt"

func TestCompareArray(){
	a := [...]int{1,2,3,4,5}
	b := [...]int{1,3,5,7,9}
	c := [...]int{1,2,3,4,6}
	d := [...]int{1,2,3,4,5}
	fmt.Println(a == b)
	// fmt.Println(a == c)	//长度不同的数组，编译出错
	fmt.Println(b == c)
	fmt.Println(a == d)		//只有当数组的长度相同，对应位置上面的元素完全相同时，数组比较才会返回True
}
const(
	Readable = 1 << iota
	Writable
	Executable
)
func TestBitClear()  {
	fmt.Println(Readable, Writable, Executable)
	a := 7	//0111
	// a = a &^ Readable	//清除掉可读性	&^ :二元操作符的操作结果是“bit clear”
							// a &^ b :表示将a中b对应位置上为1的清零
	// fmt.Println(a)	//6: 0110

	//依次清除掉可读性、可写性、可执行行
	fmt.Println(a &^ Readable == Readable, a &^ Writable == Writable, a &^ Executable == Executable)

}

func main()  {
	// TestCompareArray()
	TestBitClear()
}