package main

import(
	"fmt"
	"time"
	"sync"
)
//判断
func isCancelled(cancelChan chan struct{}) bool{
	select{
	case a := <- cancelChan:
		fmt.Println(a, "......")
		return true
	default:
		return false
	}
}
func cancel_1(cancelChan chan struct{}){
	cancelChan <- struct{}{}
}
func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)	//close(chan)是一种广播机制，所有协程都会监测到
}
func TestCancel(wg *sync.WaitGroup){
	cancelChan := make(chan struct{}, 0)
	for i:=0; i<5; i++{
		wg.Add(1)
		go func(i int, cancelCh chan struct{}){
			for{
				if isCancelled(cancelCh){
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canneled")
			wg.Done()
		}(i, cancelChan)
	}
	// cancel_1(cancelChan)	//这种方式只有 4 Canneled，其他协程没有被取消
	cancel_2(cancelChan)
	// time.Sleep(time.Millisecond * 1)
}

func main()  {
	var wg sync.WaitGroup
	TestCancel(&wg)
	wg.Wait()
	fmt.Println("Continue Running......")
}