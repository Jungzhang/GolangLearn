package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)


type DbCfg struct {
	host    string `json:"host"`
	port    int    `json:"port"`
	dbName  string `json:"db_name"`
	user    string `json:"user"`
	pass    string `json:"pass"`
	charset string `json:"charset"`
}

func main() {

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	var db DbCfg

	err = json.Unmarshal(data, &db)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(db)

}
