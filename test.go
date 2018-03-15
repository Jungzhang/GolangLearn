package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

type DbCfg struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	DbName  string `json:"db_name"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Charset string `json:"charset"`
}

func main() {

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	data1 := []byte(`{"host": "127.0.0.1","port": 3306,"db_name": "ginDemo","user": "root","pass": "","charset": "utf-8"}`)

	var db DbCfg

	var m map[string]interface{}

	err = json.Unmarshal(data1, &m)
	err = json.Unmarshal(data1, &db)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("m", m)

	fmt.Println(db)

}
