/*package main

import (
	"fmt"
	"reflect"

	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x) //
	display(name, reflect.ValueOf(x))
}

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

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s",path,v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]",path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

type movie struct {
	Title, Subtitle 	string
	Year				int
	Color               map[string]string
	Actor 				[]string
	Oscars 				[]string
	Sequel 				*string
}

func main() {
	sequel := "jianglei"
	strangelove := movie {
		Title:  			"Dr. Strangelove",
		Subtitle: 			"How i Learned to stop worrying and lov the bomb",
		Year: 				1654,
		Color:   			map[string]string {
			"Di. Stanefe": 			"pertew",
			"dwefrefsfee":			"feffdgvt",
			"drfgerfesaf":			"fvbnhyju",
		},	
		Actor: 				[]string{
            "Peter Sellers",
            "George C. Scott",
            "Sterling Hayden",
        },
		Oscars:             []string {
			"Best Actor (Nomin)",
			"frefsfd ewfrgrtg",
			"fvgnlykhy p[rde[fpor]]",
		},
		Sequel: &sequel,
	}
	Display("strangelove", strangelove)
}*/