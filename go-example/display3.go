/*package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf * bytes.Buffer, v reflect.Value) error{
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
		//buf.WriteByte('\n')

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
		//buf.WriteByte('\n')
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
			buf.WriteByte('\n')
		}
		buf.WriteByte(')')
		buf.WriteByte('\n')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
			//buf.WriteByte('\n')
		}
		buf.WriteByte(')')
		//buf.WriteByte('\n')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte,error) {
	var buf bytes.Buffer //缓冲区
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(),nil
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
	buf, err := Marshal(strangelove)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%s\n", buf)
	}

	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 64
	fmt.Println(x)
}*/