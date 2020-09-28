package main

import(
	"fmt"
	"time"
)

type IntConv func(op int)int
func TestTimeSpend(inner IntConv) IntConv {
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend is", time.Since(start).Seconds())
		return ret
	}
}

func SlowFunc(op int) int{
	// fmt.Println("slowFun is running......", op)
	time.Sleep(time.Second*1)	//睡眠一秒
	return op
}
func TestFn(){
	tsSF := TestTimeSpend(SlowFunc)	//返回的函数里面调用了SlowFunc函数，相当于对SlowFunc函数做了一层封装，使SlowFunc函数具备时间花费计算功能。
	fmt.Println(tsSF(10))
}



func main()  {
	TestFn()
}