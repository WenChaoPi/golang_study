package main

import(
	"fmt"
	"errors"
	"strconv"
)

//区分错误类型的时候，最好定义预制的错误
var LessThenTwoError  = errors.New("n should not less than 2")
var MoreThenHundredError = errors.New("n should not larger than 100")

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

func GetFibonacci2(str string)  {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil{	//尽量及早失败，避免嵌套
		fmt.Println("Error:", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(list)
}

func TestGetFibonacciByStr(){
	GetFibonacci2("10")
	GetFibonacci2("-1")
	GetFibonacci2("101")
}

func TestGetFibonacci(){
	if fiblist, err := GetFibonacci(101); err != nil{	//以后的项目中也应该写成err != nil 这样可以及早的判断错误，避免嵌套
		if err == LessThenTwoError{
			fmt.Println("It is less.")
		} else if err == MoreThenHundredError{
			fmt.Println("It is more.")
		}
		// fmt.Println(err)
		return
	}else {
		fmt.Println(fiblist)
	}

}

func main()  {
	// TestGetFibonacci()
	TestGetFibonacciByStr()
}