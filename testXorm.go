package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

type userinfo struct {
	Account            string
	Passwd             string
	Acctype            int
	SkipReasonOption   string
	Username           string
	Photo              string
	SkipReasonOption_1 string
}

func main() {
	orm, err := xorm.NewEngine("mysql", "root:@/ginDemo?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	u := new(userinfo)
	u.Username = "aaaaaaaa"
	has, err := orm.Get(u)
	if err != nil {
		log.Println(err)
	}
	if !has {
		log.Println(has)
	}
}
