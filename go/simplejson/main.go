package main 
import(
    "fmt"
    "github.com/bitly/go-simplejson"
    "log"
)

func main() {
    js, err := simplejson.NewJson([]byte(`{
            "string_array": ["asdf", "ghjk", "zxcv"],
            "string_array_null": ["abc", null, "efg"],
            "array": [1, "2", 3],
            "arraywithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "string": "simplejson",
            "bool": true,
            "sub_obj": {"a": 1}
    }`))
    if err != nil {
        log.Fatal(err)
    }
    _, ok := js.CheckGet("test") //true ,有该字段
    fmt.Println(ok)
    missJson, ok := js.Get("test").CheckGet("string_array") 
    fmt.Println(ok)//false，没有该字段
    fmt.Printf("string array\n")
    fmt.Println(missJson) //<nil>
    if content, err := missJson.MarshalJSON(); err != nil {
	fmt.Printf("MarshalJSON error = %v\n", err)
    } else {
	fmt.Printf("MarshalJSON = %s\n", content)
    }

    fmt.Println(js.Get("test").Get("string").MustString()) //simplejson
    fmt.Println(js.Get("test").Get("missing_array").MustArray([]interface{}{"1", 2, "3"})) //[1 2 3]

    msa := js.Get("test").Get("string_array").MustStringArray() 
    for _,v := range msa{
        fmt.Println(v)
    }
    //返回：
    //asdf
    //ghjk
    //zxcv

    gp, _ := js.GetPath("test", "string").String()
    fmt.Println(gp) //simplejson

    js.Set("test2", "setTest")
    fmt.Println(js.Get("test2").MustString()) //setTest
    js.Del("test2")
    fmt.Println(js.Get("test2").MustString()) //为空

    s := make([]string, 2)
    s[0] = "test2"
    s[1] = "name"
    js.SetPath(s, "testSetPath")
    gp, _ = js.GetPath("test2", "name").String()
    fmt.Println(gp) //testSetPath
}
