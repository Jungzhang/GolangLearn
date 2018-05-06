package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type test struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func main() {
	data := `{
	  "code": 0,
	  "msg": "ok",
	  "data": {
		"types": [
		  {
			"id": "1001",
			"desc": "哈哈哈"
		  },
		  {
			"id": "1002",
			"desc": "你猜"
		  },
		  {
			"id": "1003",
			"desc": "测试"
		  },
		  {
			"id": "1004",
			"desc": "This is test"
		  }
		],
		"block": 0,
		"info": "http://www.google.com"
	  }
	}`
	aaa := &test{}
	err := json.Unmarshal([]byte(data), aaa)
	fmt.Println(aaa.Data)
	if err != nil {
		log.Fatal(err)
	}
}
