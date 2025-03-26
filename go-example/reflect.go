/*package main

import (
	"fmt"
	"reflect"
)

type Dog struct{
	name string
	age int
}

func (t Dog) T1(nums []int) {
	fmt.Println(nums)
	fmt.Println("t1")
}

func (t Dog) T2() {
	fmt.Println("t2")
}

func main() {
	var tt Dog
	tt.age = 10
	tt.name = "xiaoming"
	getType := reflect.TypeOf(tt)
	getvalue := reflect.ValueOf(tt)
	fmt.Println(getvalue)
	for i := 0; i < getType.NumMethod(); i++ {
		met := getType.Method(i)
		fmt.Printf("%s, %s, %s, %d\n", met.Type, met.Type.Kind(), met.Name, met.Index)
	}
}*/