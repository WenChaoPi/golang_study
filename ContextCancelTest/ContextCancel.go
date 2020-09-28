package main

import(
	"fmt"
	"time"
	"sync"
	"context"
)
//判断
func isCancelled(ctx context.Context) bool{
	select{
	case <- ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(wg *sync.WaitGroup){
	ctx, cancel := context.WithCancel(context.Background())	//返回上下文和cancel方法
	for i:=0; i<5; i++{
		wg.Add(1)
		go func(i int, ctx context.Context){
			for{
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canneled")
			wg.Done()
		}(i, ctx)
	}
	cancel()
}

func main()  {
	var wg sync.WaitGroup
	TestCancel(&wg)
	wg.Wait()
	fmt.Println("Continue Running......")
}