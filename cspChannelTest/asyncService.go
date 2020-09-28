package main

import(
	"fmt"
	"time"
)
func Service() string {
	time.Sleep(time.Second * 1)
	return "Done"
}
func OtherTask(){
	fmt.Println("working on something else")
	time.Sleep(time.Second * 2)
	fmt.Println("Task is done.")
}
func TestService(){
	fmt.Println(Service())
	OtherTask()
	//以上是串行执行的
}

//将串行改成异步
func AsyncService() chan string{
	// retCh := make(chan string)	//声明一个chan 存放string 不带buffer
	retCh := make(chan string, 1)	//定义buffer channel	不阻塞service，效率更高
	go func (){
		ret := Service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}
func TestAsyncService(){	//异步service
	retCh := AsyncService()
	OtherTask()
	fmt.Println(<-retCh)	//这里会阻塞等待channel里面的内容
	time.Sleep(time.Second * 1)	//service exited.会最后输出，retCh信道会导致上面的协程阻塞，应该改成retCh := make(chan string, 1)
}

func main()  {
	// TestService()	//同步
	TestAsyncService()
}