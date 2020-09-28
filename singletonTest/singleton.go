package main

import(
	"fmt"
	"sync"
	"unsafe"
)
type Singleton struct{

}
var singletonInstance *Singleton
var once sync.Once
func GetSingletonInstance() *Singleton{
	once.Do(func(){		//确保无论创建多少个协程，这里只会创建一个Singleton实例
		fmt.Println("create obj")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}
func TestGetSingleton(){
	var wg sync.WaitGroup
	for i:=0; i<5; i++{
		wg.Add(1)
		go func(){
			obj := GetSingletonInstance()
			fmt.Printf("%x\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

func main()  {
	TestGetSingleton()
}