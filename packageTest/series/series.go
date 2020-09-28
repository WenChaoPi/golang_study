package series

import(
	"fmt"
	"errors"
)

func init(){
	fmt.Println("init1")
}
func init(){
	fmt.Println("init2")
}

var LessThenTwoError = errors.New("less")
var MoreThenHundredError = errors.New("more")
func GetFibonacci(n int) ([]int, error) {	//多返回值，切片和error
	if n < 2{
		return nil, LessThenTwoError	//自定义错误使用errors.New()
	}
	if n > 100{
		return nil, MoreThenHundredError
	}
	fibList := []int{1,1}
	for i:=2; i<n; i++{
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}


func square(x, y int) int{	//小写字母开头的方法，包外无法访问。
	return x * y
}

func Square(x, y int) int{	//小写字母开头的方法，包外无法访问。
	return x * y
}
