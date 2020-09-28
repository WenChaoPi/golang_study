package main

import(
	"fmt"
	"unsafe"
)

type Employee struct{
	Id	string
	Name	string
	Age	int
}
func TestCreateEmployeeObj(){
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)		//new返回的是引用/指针	仍然使用 . 获取里面的成员变量
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 18
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Printf("e is %T\ne2 is %T", e, e2)
}

//给对象定义行为
//这种方式方法调用时，结构体实例被复制了一份，
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age-%d",e.Id, e.Name, e.Age)
}

//推荐使用这种方法，避免了内存拷贝； 使用的是同一个结构体实例
// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age-%d",e.Id, e.Name, e.Age)
// }
func TestStructOperations()  {
	e := Employee{Id:"1", Name:"Tim", Age:19}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	fmt.Println(e.String())
}

func main()  {
	TestCreateEmployeeObj()
	// TestStructOperations()
}