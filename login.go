package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method =", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username =", r.Form["username"])
		fmt.Println("password =", r.Form["passwd"])
	}
}

func main() {
	http.HandleFunc("/login", handlerLogin)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
