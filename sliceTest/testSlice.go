package main

import "fmt"

func TestSliceInit()  {
	var s0 []int 	//定义切片
	fmt.Println(len(s0), cap(s0))
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 3, 5)		//length:表示可以访问的元素的个数，capacity:表示可以填充的元素个数
	fmt.Println(len(s2), cap(s2))
	// fmt.Println(s2[0], s2[1], s2[2], s2[3], s2[4])	//越界异常
	fmt.Println(s2[0], s2[1], s2[2])
	s2  = append(s2, 1)
	fmt.Println(s2[0], s2[1], s2[2], s2[3])
}
func TestSliceGrowing()  {	//slice扩容时每次扩大2倍
	s := []int{}
	for i:=0; i<10; i++{
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}
/**
1 1
2 2
3 4
4 4
5 8
6 8
7 8
8 8
9 16
10 16
*/

func TestSliceShareMemory()  {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Println(len(year), cap(year))
	spring := year[:3]
	fmt.Println(spring, len(spring), cap(spring))

	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))

	summer[0] = "Unknown"	//由于切片共享存储结构，所以修改其中的值，会全局影响
	fmt.Println(spring)

	
}

func TestSliceComparing()  {
	// a := []int{1,2,3}
	// b := []int{1,2,3}
	// fmt.Println(a == b)	//(slice can only be compared to nil)

}

func main()  {
	// TestSliceInit()
	// TestSliceGrowing()
	TestSliceShareMemory()
}