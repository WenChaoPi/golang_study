package main

import(
	"fmt"
	"sync"
)

func dataProducter(ch chan int, wg *sync.WaitGroup)  {
	go func(){
		for i:=0; i<10; i++{
			ch <- i
		}
		close(ch)
		// ch <- 11	//往已经关闭的新到上面发送数据，会报错panic: send on closed channel
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup)  {
	go func(){
		for {	
			if data,ok := <- ch; ok{//false代表通道关闭；true代表正常接收
				fmt.Println(data)
			}else {
				fmt.Println("channel is closed.")
				break
			}
		}
		wg.Done()
	}()
}
//问题：接收者如何知道发送者发送完数据了（添加特殊token）；发送者如何知道有多少接收者，根据接收者个数添加多少token
//解决方案：发送者发送完数据关闭channel
//接收者就根据就收到的ok判断是否通道关闭
func TestCloseChannel()  {
	var wg sync.WaitGroup
	ch := make(chan int, 10)	//如果有容量会如何
	wg.Add(1)
	dataProducter(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
func main()  {
	TestCloseChannel()
}