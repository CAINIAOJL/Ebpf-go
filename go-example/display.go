/*package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Any(value interface{}) string {
	//a := reflect.TypeOf(value)
	return formatAtom(reflect.ValueOf(value))
}

//value中包含着type
func formatAtom(v reflect.Value) string{
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int16,reflect.Int8,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint16,reflect.Uint8,reflect.Uint32,reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice,reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16) //十六进制
	default: //Array,Struct,Interface
		return v.Type().String() + " Value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	ch := make(chan <- int, 2)
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any(ch))
}*/