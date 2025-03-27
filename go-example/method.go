package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func Print(v interface{}) {
	x := reflect.ValueOf(v)
	y := reflect.TypeOf(v)
	t := x.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < x.NumMethod(); i++ {
		methodtype := x.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t,t.Method(i).Name, strings.TrimPrefix(methodtype.String(), "func")) //删除“func”输出函数名
	}
	fmt.Println()
	for i := 0; i < y.NumMethod(); i++ {
		methodtype := y.Method(i).Type
		fmt.Printf("func (%s) %s%s\n", y,y.Method(i).Name, strings.TrimPrefix(methodtype.String(), "func")) //删除“func”输出函数名
	}
	
}

func main() {
	Print(time.Hour)
	a := make([]int, 2)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
}