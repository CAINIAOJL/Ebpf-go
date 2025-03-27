/*package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func search(req *http.Request, resp http.ResponseWriter) {
	var data struct {
		Labels      	[]string
		MaxResults 		int
		Exact 			bool
	}
	data.MaxResults = 10
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return 
	}
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldinfo := v.Type().Field(i)
		tag := fieldinfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldinfo.Name)
		}
		fields[name] = v.Field(i)
	}

	for name, values := range req.Form {
		f := fields[name] //reflectValue
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}*/

/*package main

import (
    "fmt"
    "net/http"
    "reflect"
    "strconv"
    "strings"
)

func search(resp http.ResponseWriter, req *http.Request) {
    var data struct {
        Labels       []string
        MaxResults   int
        Exact        bool
    }
    data.MaxResults = 10
    if err := Unpack(req, &data); err != nil {
        http.Error(resp, err.Error(), http.StatusBadRequest)
        return
    }
    // 打印解析后的数据
    fmt.Fprintf(resp, "Labels: %v\n", data.Labels)
    fmt.Fprintf(resp, "MaxResults: %d\n", data.MaxResults)
    fmt.Fprintf(resp, "Exact: %v\n", data.Exact)
}

func Unpack(req *http.Request, ptr interface{}) error {
    if err := req.ParseForm(); err != nil {
        return err
    }

    fields := make(map[string]reflect.Value)
    v := reflect.ValueOf(ptr).Elem()
    for i := 0; i < v.NumField(); i++ {
        fieldinfo := v.Type().Field(i)
        tag := fieldinfo.Tag
        name := tag.Get("http")
        if name == "" {
            name = strings.ToLower(fieldinfo.Name)
        }
        fields[name] = v.Field(i)
    }

    for name, values := range req.Form {
        f := fields[name]
        if !f.IsValid() {
            continue
        }
        for _, value := range values {
            if f.Kind() == reflect.Slice {
                elem := reflect.New(f.Type().Elem()).Elem()
                if err := populate(elem, value); err != nil {
                    return fmt.Errorf("%s: %v", name, err)
                }
                f.Set(reflect.Append(f, elem))
            } else {
                if err := populate(f, value); err != nil {
                    return fmt.Errorf("%s: %v", name, err)
                }
            }
        }
    }
    return nil
}

func populate(v reflect.Value, value string) error {
    switch v.Kind() {
    case reflect.String:
        v.SetString(value)
    case reflect.Int:
        i, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return err
        }
        v.SetInt(i)
    case reflect.Bool:
        b, err := strconv.ParseBool(value)
        if err != nil {
            return err
        }
        v.SetBool(b)
    default:
        return fmt.Errorf("unsupported kind %s", v.Type())
    }
    return nil
}

func main() {
    http.HandleFunc("/search", search)
    fmt.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}*/
