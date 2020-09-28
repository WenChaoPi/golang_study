package main

import(
	"fmt"
	"time"
)

func TestGroutine(){
	for i:=0; i<10;i++{
		go func (i int){
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 1000)
}

func main()  {
	TestGroutine()
}