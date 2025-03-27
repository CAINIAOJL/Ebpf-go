/*package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scan scanner.Scanner
	token rune //标记
}

func (lex *lexer) next() { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }


func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value)  {
	if lex.text() == "nil" {
		v.Set(reflect.Zero(v.Type()))
		lex.next()
		return
	}

	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		if v.Kind() == reflect.Ptr {
            v = v.Elem()
        }
		v.SetString(s)
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//没有遇到EOF，')'
func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

//((Title "Dr. Strangelove") (Subtitle "How i Learned to stop worrying and lov the bomb") (Year 1654) (Color (("dwefrefsfee" "feffdgvt") ("drfgerfesaf" "fvbnhyju") ("Di. Stanefe" "pertew"))) (Actor ("Peter Sellers" "George C. Scott" "Sterling Hayden")) (Oscars ("Best Actor (Nomin)" "frefsfd ewfrgrtg" "fvgnlykhy p[rde[fpor]]")) (Sequel "jianglei"))

func readList(lex *lexer, v reflect.Value)  {
	switch v.Kind() {
	case reflect.Array: //(item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		} 
	case reflect.Slice: //(item ...)
		for !endList(lex) {
			elem := reflect.New(v.Type().Elem()).Elem()
			read(lex, elem)
			v.Set(reflect.Append(v, elem))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(') //消耗'('
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')') //消耗')'
		}
	case reflect.Interface:
        // 简单处理，这里假设将interface{}当作指针类型处理
        if v.IsNil() {
            v.Set(reflect.New(v.Type().Elem()))
        }
        readList(lex, v.Elem())
	default:
		panic(fmt.Sprintf("can not decode list into %v", v.Type()))
	}
}

func unmarshal(data []byte, v interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data)) //从[]byte data中读取数据到scanner中
	lex.next()
	defer func ()  {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}	
	}()
	value := reflect.ValueOf(v).Elem()
	if value.Kind() != reflect.Ptr {
		value = reflect.ValueOf(&v).Elem()
	}
	read(lex, value)
	return nil
}

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
			//buf.WriteByte('\n')
		}
		buf.WriteByte(')')
		//buf.WriteByte('\n')
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
			"fvgnlykhy prdefpor",
		},
		Sequel: &sequel,
	}
	buf, err := Marshal(strangelove)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Marshai begin:")
		fmt.Printf("%s\n", buf)
		fmt.Println("Marshai end:")
	}

	if err = unmarshal(buf, &strangelove); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("UnMarshai begin:")
		fmt.Printf("%+v\n", strangelove)
		fmt.Println("UnMarshai end:")

	}
}

/*package main

import (
    "bytes"
    "fmt"
    "reflect"
    "strconv"
    "text/scanner"
)

type lexer struct {
    scan scanner.Scanner
    token rune //标记
}

func (lex *lexer) next() { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
    if lex.token != want {
        panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
    }
    lex.next()
}

func read(lex *lexer, v reflect.Value) {
    if lex.text() == "nil" {
        v.Set(reflect.Zero(v.Type()))
        lex.next()
        return
    }

    switch lex.token {
    case scanner.Ident:
        if lex.text() == "nil" {
            v.Set(reflect.Zero(v.Type()))
            lex.next()
            return
        }
    case scanner.Int:
        i, _ := strconv.Atoi(lex.text())
        v.SetInt(int64(i))
        lex.next()
        return
    case '(':
        lex.next()
        readList(lex, v)
        lex.next()
        return
    }
    panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//没有遇到EOF，')'
func endList(lex *lexer) bool {
    switch lex.token {
    case scanner.EOF:
        panic("end of file")
    case ')':
        return true
    }
    return false
}

//((Title "Dr. Strangelove") (Subtitle "How i Learned to stop worrying and lov the bomb") (Year 1654) (Color (("dwefrefsfee" "feffdgvt") ("drfgerfesaf" "fvbnhyju") ("Di. Stanefe" "pertew"))) (Actor ("Peter Sellers" "George C. Scott" "Sterling Hayden")) (Oscars ("Best Actor (Nomin)" "frefsfd ewfrgrtg" "fvgnlykhy p[rde[fpor]]")) (Sequel "jianglei"))

func readList(lex *lexer, v reflect.Value) {
    switch v.Kind() {
    case reflect.Array: //(item ...)
        for i := 0;!endList(lex); i++ {
            read(lex, v.Index(i))
        }
    case reflect.Slice: //(item ...)
        for!endList(lex) {
            elem := reflect.New(v.Type().Elem()).Elem()
            read(lex, elem)
            v.Set(reflect.Append(v, elem))
        }
    case reflect.Struct:
        for!endList(lex) {
            lex.consume('(') //消耗'('
            if lex.token != scanner.Ident {
                panic(fmt.Sprintf("got token %q, want field name", lex.text()))
            }
            name := lex.text()
            lex.next()
            read(lex, v.FieldByName(name))
            lex.consume(')')
        }
    case reflect.Map:
        v.Set(reflect.MakeMap(v.Type()))
        for!endList(lex) {
            lex.consume('(')
            key := reflect.New(v.Type().Key()).Elem()
            read(lex, key)
            value := reflect.New(v.Type().Elem()).Elem()
            read(lex, value)
            v.SetMapIndex(key, value)
            lex.consume(')') //消耗')'
        }
    default:
        panic(fmt.Sprintf("can not decode list into %v", v.Type()))
    }
}

func unmarshal(data []byte, v interface{}) (err error) {
    lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
    lex.scan.Init(bytes.NewReader(data)) //从[]byte data中读取数据到scanner中
    lex.next()
    defer func() {
        if x := recover(); x != nil {
            err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
        }
    }()

    // 确保传入的是指针类型
    value := reflect.ValueOf(v)
    if value.Kind() != reflect.Ptr {
        value = reflect.ValueOf(&v).Elem()
    }
    read(lex, value)
    return nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
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
            //buf.WriteByte('\n')
        }
        buf.WriteByte(')')
        //buf.WriteByte('\n')
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

func Marshal(v interface{}) ([]byte, error) {
    var buf bytes.Buffer //缓冲区
    if err := encode(&buf, reflect.ValueOf(v)); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

type movie struct {
    Title, Subtitle string
    Year            int
    Color           map[string]string
    Actor           []string
    Oscars          []string
    Sequel          *string
}

func main() {
    sequel := "jianglei"
    strangelove := movie{
        Title:    "Dr. Strangelove",
        Subtitle: "How i Learned to stop worrying and lov the bomb",
        Year:     1654,
        Color: map[string]string{
            "Di. Stanefe":  "pertew",
            "dwefrefsfee":  "feffdgvt",
            "drfgerfesaf": "fvbnhyju",
        },
        Actor: []string{
            "Peter Sellers",
            "George C. Scott",
            "Sterling Hayden",
        },
        Oscars: []string{
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
        fmt.Println("Marshai begin:")
        fmt.Printf("%s\n", buf)
        fmt.Println("Marshai end:")
    }

    // 注意这里传入的是指针
    if err = unmarshal(buf, &strangelove); err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("UnMarshai begin:")
        fmt.Printf("%+v\n", strangelove)
        fmt.Println("UnMarshai end:")
    }
}*/