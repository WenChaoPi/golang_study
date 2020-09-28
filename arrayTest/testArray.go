package main

import "fmt"

func TestArrayInit()  {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,2,3,4,5}
	fmt.Println(arr[1], arr[2])
	fmt.Println(arr1, arr3)
}
func TestArrayTravel(){
	arr3 := [...]int{1,2,3,4,5}
	for i:=1; i<len(arr3); i++{
		fmt.Println(arr3[i])
	}
	for idx, ele := range arr3{
		fmt.Println(idx, ele)
	}

	for _, ele := range arr3{	//使用下划线占位
		fmt.Println(ele)
	}
}
func TestArraySection()  {
	arr3 := [...]int{1,2,3,4,5}
	a := arr3[:3]
	b := arr3[:]
	c := arr3[2:]
	fmt.Println(a,b,c)
}

func main()  {
	// TestArrayInit()
	// TestArrayTravel()
	TestArraySection()
}