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

func AllResponse() string {
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
	//等待所有任务执行完成，返回
	finalRet := ""
	for j:=0; j<numOfRunner; j++{
		finalRet += <-ch + "\n"
	}
	return finalRet
	//也可以使用之前的wg.Wait()
}

func main()  {
	fmt.Println("Before: ", runtime.NumGoroutine())	
	fmt.Println(AllResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After: ", runtime.NumGoroutine())
}