package main

import(
	"fmt"
	// "os"
	"errors"
)



func main()  {
	// defer func (){
	// 	fmt.Println("Finally")
	// }()

	defer func (){	//匿名自执行函数
		if err := recover(); err != nil{
			fmt.Println("revover from ", err)	//在错误恢复时只记录错误日志是不好的习惯，有可能是系统资源消耗完，此进程仍然无法提供服务。
												//应该尽量处理错误
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Somthing wrong!"))	//调用defer指定函数；输出调用栈信息

	//exit status 255
	// os.Exit(-1)	//os.Exit(-1)不调用defer指定函数；不输出当前调用栈信息
}