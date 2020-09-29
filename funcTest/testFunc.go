package main

import (
	"fmt"
	"math/rand"
	"time"
)


func returnMultiValues() (int, int){
	rand.Seed(time.Now().UnixNano())	//设置随机的种子，要不然每次产生的随机数都是一样的
	return rand.Intn(10), rand.Intn(20)
}
func TestTimeSpend(inner func(op int)int) func (op int)int {
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


//可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops{
		ret += op
	}
	return ret
}
func TestVarParam(){
	fmt.Println(Sum(1,2,3,4))
	fmt.Println(Sum(1,2,3,4,5))
}


//defer函数
func Clear() {
	fmt.Println("Clear resources.")
}
func TestDefer() {
	defer Clear()
	fmt.Println("start")
	panic("err")	//即便是发生了panic不可修复的异常，defer函数仍然会执行
	// fmt.Println("afer panic is not ok")		//panic之后的飞defer代码不可执行（不可达）
}

func main()  {
	// a,b := returnMultiValues()
	// fmt.Println(a,b)
	// TestFn()
	// TestVarParam()
	TestDefer()
}