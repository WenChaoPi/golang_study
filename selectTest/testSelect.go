package main

import(
	"fmt"
	"time"
)
func Service() string{
	time.Sleep(time.Millisecond * 50)
	return "Service Done."
}
func AsyncService() chan string{
	retCh := make(chan string, 1)
	go func(){
		ret := Service();
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("Service exited.")
	}()
	return retCh
}
func TestSelect(){
	select{
	case ret := <- AsyncService():
		fmt.Println(ret)
	case <- time.After(time.Millisecond * 100):	//超时控制；多渠道选择
		fmt.Println("time out...")
	}
}


func main()  {
	TestSelect()
}