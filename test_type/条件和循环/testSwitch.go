package main
import "fmt"

func TestSwitch()  {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Println("Even")
		case 1, 3:
			fmt.Println("Odd")
		default:
			fmt.Println("is not 1——3")
		}
	}
}

func TestSwitchCaseContaion()  {
	for i := 0; i < 5; i++ {
		switch {
		case i%2==0:
			fmt.Println("Even")
		case i%2==1:
			fmt.Println("Odd")
		default:
			fmt.Println("is not 1——3")
		}
	}
}

func main()  {
	TestSwitchCaseContaion()
}