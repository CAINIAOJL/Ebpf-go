/*package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//以下的``中包含的是一种解析数据的格式,是一一对应的映射关系
//比如：Address 对应ini文件内的address字段
//在进行反射的时候将address对应的数据解析到Address字段中。

type MysqlConfig struct {
	Address    string `ini:"address"`
	Port       int    `ini:"port"`
	Username   string `ini:"username"`
	Userpasswd string `ini:"userpasswd"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redies"`
	// int         `ini:"mysql"`
}

func (c Config) Println() {
	mType := reflect.TypeOf(c)
	mValue := reflect.ValueOf(c)
	for i := 0; i < mType.NumField(); i++ {
		fmt.Println("成员结构体属性如下:", mType.Field(i).Name)
		tempValue := mValue.FieldByName(mType.Field(i).Name)
		// tempValue := mValue.FieldByName(mType.Field(i).Name)
		// 获取到的是值域,可以通过该对象进行每个字段的值进行获取
		tempType := tempValue.Type()
		// tempType := tempValue.Type()
		//获取到的是值域的具体类型信息

		// 再次迭代取值
		for j := 0; j < tempType.NumField(); j++ {
			fmt.Println("属性类型：", tempType.Field(j).Type, "属性名：", tempType.Field(j).Name, "属性值：", tempValue.Field(j))
		}
		fmt.Println()
	}
}

func Loading_INI(filename string, data interface{}) error {

	// 获取data的类型
	dataType := reflect.TypeOf(data)

	// 获取data对应的值
	dataValue := reflect.ValueOf(data)

	// 判断传进来的参数是不是指针类型,并且判断一下这个指针类型指向的是不是结构体
	if dataType.Kind() != reflect.Ptr || dataType.Elem().Kind() != reflect.Struct {
		return errors.New("the 2 args must be struct ptr")
	}

	// 读取配置文件
	bytes, err := ioutil.ReadFile("myconfig.ini")
	if err != nil {
		return err
	}
	iniLines := strings.Split(string(bytes), "\r\n")

	// 用于暂时存储ini配置文件内的模块名对应的结构体名
	var ComfigModelStructName string

	// 迭代每一行数据,进行对应的操作
	for index, line := range iniLines {

		// 去除行前面的空格
		line = strings.TrimSpace(line)
		// 判断是不是注释行与空行,是的话直接跳过本次循环
		if len(line) == 0 || (line[0] == '#' || line[0] == ';') { //排除掉空行,以及注释
			continue
		} else if line[0] == '[' {
			// 如果该行没有]
			if line[len(line)-1] != ']' {
				return fmt.Errorf("the '[' has nomatch ']' in %d line", index)
				//如果该行里面有[]却没有字段
			} else if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				return fmt.Errorf("the [ ] has none in %d line", index)
				//字段书写规范
			} else {
				// 将配置文件内的模块名拿出来
				for i := 0; i < dataType.Elem().NumField(); i++ {
					if line[1:len(line)-1] == string(dataType.Elem().Field(i).Tag.Get("ini")) {
						ComfigModelStructName = dataType.Elem().Field(i).Name
					}
				}
				fmt.Println("结构体", ComfigModelStructName, "正在解析!")
			}
		} else {
			// 将数据以等号为分隔符进行分割
			tempIndex := strings.Index(line, "=")
			Mkey := line[0:tempIndex]
			Mvalue := line[tempIndex+1:]
			// 如果只有值没有键，那么抛出异常,只有键没有值，那么数据为默认值直跳过本次循环
			if len(strings.TrimSpace(Mkey)) == 0 {
				return fmt.Errorf("the Key is null in %d line", index+1)
			} else if len(strings.TrimSpace(Mvalue)) == 0 {
				continue
			} else {
				// 根据对应的结构体名拿到结构体具体信息
				sValue := dataValue.Elem().FieldByName(ComfigModelStructName)
				sType := sValue.Type()
				if sType.Kind() != reflect.Struct {
					return errors.New("the data members must be Struct")
				}
				for i := 0; i < sType.NumField(); i++ {
					if sType.Field(i).Tag.Get("ini") == Mkey {
						switch sType.Field(i).Type.Kind() {
						case reflect.String:
							sValue.Field(i).SetString(Mvalue)
						case reflect.Bool:
							mbool, err := strconv.ParseBool(Mvalue)
							if err != nil {
								return fmt.Errorf("the value must be true or false in %d line", index+1)
							}
							sValue.Field(i).SetBool(mbool)
						case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
							mint, err := strconv.ParseInt(Mvalue, 10, 64)
							if err != nil {
								return fmt.Errorf("the value must be number in %d line", index+1)
							}
							sValue.Field(i).SetInt(mint)
						case reflect.Float32, reflect.Float64:
							mfloat, err := strconv.ParseInt(Mvalue, 10, 64)
							if err != nil {
								return fmt.Errorf("the value must be number in %d line", index+1)
							}
							sValue.Field(i).SetInt(mfloat)
						}
					}
				}
			}
		}
	}
	fmt.Println("所有配置项解析完毕!")
	return nil
}

/*
dataType.Elem().Kind()		写法会有很大的隐患，因为如果dataType不是指针类型就会报错，Elem方法只是为了拿到指针指向的元素
if dataType.Kind() != reflect.Ptr || dataType.Elem().Kind() != reflect.Struct 不报错
因为go语言判断的时候有阻断机制,在第一个判断条件为为真时不会去判断第二个条件,所以||后面的语句没有执行,不会报错
*/
/*func main() {

	var cg1 Config
	err := Loading_INI("myconfig.ini", &cg1)
	if err != nil {
		fmt.Println("error : ", err)
	}
	cg1.Println()
}*/

