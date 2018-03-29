package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Rsp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func main() {
	jjj := `{"data": {
        "skip_types": [
            {
                "id": 1001,
                "desc": "6666"
            },
            {
                "id": 1002,
                "desc": "kkkkk"
            },
            {
                "id": 1003,
                "desc": "xixixixi"
            },
            {
                "id": 1004,
                "desc": "hahah"
            }
        ],
        "block": "true",
        "info": "http:\/\/www.google.com"
    },
    "code": 0,
    "msg": "ok"
}`
	rrr := &Rsp{}
	jj := []byte(jjj)

	err := json.Unmarshal(jj, rrr)
	if err != nil {
		log.Println(err)
	}
	reasons := rrr.Data["skip_types"]
	for _, resaon := range reasons.([]interface{}) {
		r := resaon.(map[string]interface{})["desc"]
		fmt.Println(r)
	}

	fmt.Println(strconv.ParseBool("1"))
}
