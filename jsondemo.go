//package main
//
//import (
//	"encoding/json"
//	"log"
//	"os"
//)
//
//func main() {
//	dec := json.NewDecoder(os.Stdin)
//	enc := json.NewEncoder(os.Stdout)
//	for {
//		var v map[string]interface{}
//		if err := dec.Decode(&v); err != nil {
//			log.Println(err)
//			return
//		}
//		for k := range v {
//			if k != "Title" {
//				v[k] = nil
//			}
//		}
//		if err := enc.Encode(&v); err != nil {
//			log.Println(err)
//		}
//	}
//}
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type School struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func main() {
	school := School{
		Name:     "某某学校",
		Location: "市区",
	}

	str, err := json.Marshal(school)

	if err != nil {
		fmt.Printf("json序列化失败%v", err)
	}

	fmt.Println(str)                 //返回结果[123 34 110 97 109 101 34 58 34 230 159 144 230 159 144 229 173 166 230 160 161 34 44 34 108 111 99 97 116 105 111 110 34 58 34 229 184 130 229 140 186 34 125]
	fmt.Println(reflect.TypeOf(str)) //返回结果[]uint8
	fmt.Println(string(str))         //返回结果{"name":"某某学校","location":"市区"}
}
