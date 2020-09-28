package main

import "fmt"

func TestMapWithFuncValue(){
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op * op}
	m[3] = func(op int) int {return op * op * op}
	fmt.Println(m[1](2), m[2](2), m[3](2))

}

func TestMapForSet()  {	//golang 中不存在set， 但可以使用map[type]bool	实现set的功能
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {//set中是否存在某个元素
		fmt.Println("1 is existing")
	} else {
		fmt.Println("1 is not existing")
	}
	mySet[3] = true
	//set的大小
	fmt.Println(len(mySet))

	//删除set中的指定的元素
	delete(mySet, 1)	//删除map中的指定key即可
	fmt.Println(len(mySet))
}

func main()  {
	// TestMapWithFuncValue()
	TestMapForSet()
}
