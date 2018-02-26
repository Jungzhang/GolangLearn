package main

import (
	"fmt"
	"io"
	"os"
)

func testErr(id int) (int, error) {

	if id < 0 {
		return 0, fmt.Errorf("id is illegal, id =%d", id)
	}
	return id, nil
}

func main() {

	var w io.Writer
	w = os.Stdout

	_, err := testErr(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	_, err = testErr(-1)
	if err != nil {
		fmt.Println(err)
	}

	v, ok := w.(*os.File)
	//v, ok := w.(*bytes.Buffer)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("断言失败")
	}

}
