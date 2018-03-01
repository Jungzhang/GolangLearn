package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
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

// 文件上传示例
func handlerFileUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10)) // 第二个参数表示进制取值为2-32之间
		token := fmt.Sprintf("%x", h.Sum(nil))
		t.Execute(w, token)
	} else if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0664)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} else {
		log.Println("Cannot identify method, method =", r.Method)
	}

}

func main() {
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/upload", handlerFileUpload)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
