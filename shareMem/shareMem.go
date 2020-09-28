package main

import(
	"fmt"
	// "time"
	"sync"
)

// func TestCounter()  {
// 	counter := 0	//此时counter在5000个协程中是共享的，存在安全性问题。
// 	for i:=0; i<5000; i++{
// 		go func(){
// 			counter++
// 		}()
// 	}
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("counter = %d", counter)
// }

// func TestCounterThreadSafe()  {
// 	var mut sync.Mutex	//定义锁
// 	counter := 0
// 	for i:=0; i<5000; i++{
// 		go func(){
// 			defer func (){
// 				mut.Unlock()	//类似于java中在final中释放锁
// 			}()
// 			mut.Lock()	//加锁
// 			counter++
// 		}()
// 	}
// 	time.Sleep(1 * time.Second)	//防止外面的协程执行太快，导致协程来不及执行
// 	fmt.Printf("counter = %d", counter)
// }

func TestCounterWaitGroup()  {
	var mut sync.Mutex	//定义锁
	var wg sync.WaitGroup
	counter := 0
	for i:=0; i<5000; i++{
		wg.Add(1)
		go func(){
			defer func (){
				mut.Unlock()	//类似于java中在final中释放锁
			}()
			mut.Lock()	//加锁
			counter++
			wg.Done()
		}()
	}
	// time.Sleep(1 * time.Second)	//防止外面的协程执行太快，导致协程来不及执行
	wg.Wait()
	fmt.Printf("counter = %d", counter)
}

func main()  {
	TestCounterWaitGroup()
}