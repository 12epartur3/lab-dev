package main

import(
	"fmt"
	"encoding/json"
	"reflect"
)

type Config struct {
	Name string
	Index IndexSet 
}

type IndexSet struct {
	set map[int64] struct{}
}

func (is *IndexSet) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		return err
	}
	valueType := reflect.TypeOf(v)
	if valueType.Kind() != reflect.Slice || valueType.Elem().Kind() != reflect.Interface {
		fmt.Printf("type error b = %+v, type = %v\n", b, reflect.TypeOf(v))
		return nil
	}
	fmt.Printf("b = %+v, type = %v\n", b, reflect.TypeOf(v))
	if _, ok := v.([]interface{}); !ok {
		fmt.Printf("err type 1\n")
		return nil
	}
	arr, _ := v.([]interface{})
	is.set = make(map[int64] struct{})
	for _, v := range arr {
		if value, ok := v.(float64); !ok {
			fmt.Printf("err type 2, type = %v\n", reflect.TypeOf(v))
			continue
		} else {
			fmt.Printf("get value %v\n", value)
			if _, exist := is.set[int64(value)]; !exist {
				is.set[int64(value)] = struct{}{}
			}
		}
	}
	return nil
}
func main() {
	//s := `{"Config":{"Name":"yytest","Index":[1,2,3]}}`
	//s := `{"Name":"yytest","Index":[1,2,3]}`
	s := `{"Index": [10, 1, 3, 3]}`
	config := Config{}
	//var config interface{}
	err := json.Unmarshal([]byte(s), &config)
	if err != nil {
		fmt.Printf("err = %v\n", err)
		return
	}
	fmt.Printf("config = %+v, type = %v\n", config, reflect.TypeOf(config))
}
