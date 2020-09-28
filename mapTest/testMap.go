package main

import "fmt"

func TestInitMap()  {
	m1 := map[int]int{1:1, 2:4, 3:8}	//map[keyType]valueType{init elements}
	for i:=1; i<10; i++{	//map扩容时，逐渐增加1; 而slice扩容时每次扩大2倍
		m1[i] = i
	}
	fmt.Println(m1)
	fmt.Println(len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	fmt.Println("len m2", len(m2))

	m3 := make(map[int]int, 10)	//使用make关键字定义map变量时，不指定length值，因为无法为map初始化零值
	fmt.Println("len m3", len(m3))
	// fmt.Println("cap m3", cap(m3))	//cap无法用于求map的cap
}

func TestAccessNotExistingKey()  {
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[1] = 0;
	fmt.Println(m1[1])
	//上面输出结果均为0，那我们如何判断m1[1]时默认的空值，还是我们设置的0

	//解决：
	if v, ok := m1[2]; ok{
		fmt.Println("key 2 is existing", v)
	} else{
		fmt.Println("key 2 is not existing", v)
	}

}
func TestTravelMap()  {
	m1 := map[int]int{1:2, 2:4, 3:8}
	for k, v := range m1{	//map的遍历是不按照顺序的
		fmt.Println(k, v)
	}
}

func main()  {
	TestInitMap()
	TestAccessNotExistingKey()
	TestTravelMap()
}