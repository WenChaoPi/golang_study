package main

import (
	"fmt"
	"reflect"
)
func CheckType(v interface{}){
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Println("unknown", t)
	}
}
func TestBasicType(){
	var f float64 = 12
	CheckType(f)
}

func main()  {
	TestBasicType()
}