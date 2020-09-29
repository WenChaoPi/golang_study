package main

import (
	"fmt"
	"time"
	"runtime"
)
//执行当前任务
func runTask(id int) string  {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result if from %d", id)
}

//问题：当第一个消息被接收（消费后），其他协程往channel里面放数据时会阻塞，（因为没有人去消费），
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string)// 因为这里使用当时不带缓存当channel，所以当第一个消息被消息时，
						   // 后面没有再次消费消息，会导致其他协程阻塞往channel里面放数据，
						   // 所以应该使用ch := make(chan string, numOfRunner)，解决了协程泄漏的问题
	for i:=0; i<numOfRunner; i++ {
		go func(i int){
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <- ch
}

func main()  {
	fmt.Println("Before: ", runtime.NumGoroutine())	//有协程泄漏问题
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After: ", runtime.NumGoroutine())
}