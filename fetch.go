package main

import (
	"os"
	"net/http"
	"fmt"
//	"io/ioutil"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		// 使用Copy复制到stdout
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Println("HTTP Code : ", resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		//fmt.Printf("%s", b)
	}
}
