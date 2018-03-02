package main

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"time"
)


func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db, err := sql.Open("mysql", "root:Jung@/learngolang?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
	if err != nil {
		log.Fatal(err)
	}

	str := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec("RaidenZhang", "rd", str)
	if err != nil {
		log.Println(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("id =", id)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err = stmt.Exec("JungZhang", id)

	if err != nil {
		log.Println(err)
	}

	// 影响的条数
	affect, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("affect =", affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var id int
		var username string
		var department string
		var created string
		err = rows.Scan(&id, &username, &department, &created)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id = %v, username = %v, department = %v, created = %v\n", id, username, department, created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}

	affect, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("affect =", affect)
}
