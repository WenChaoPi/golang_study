package main
import(
	"golang_study/packageTest/series"
	"fmt"
	// "github.com/easierway/concurrent_map"
)
func TestPackage()  {
	var(
		FibList []int
		err error
	)
	if FibList, err = series.GetFibonacci(10); err != nil{
		fmt.Println(err)
		return
	} 
	fmt.Println(FibList)

	// fmt.Println(series.square(10,9))	//无法访问

	fmt.Println(series.Square(10,9))

	//尽管这里两次调用series包中函数，但series包中但init函数只执行一次。
	//一个包中可以包含多个init函数，但其他非init函数不可以重名
}

func main()  {
	TestPackage()
}