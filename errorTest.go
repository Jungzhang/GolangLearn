package main

import "fmt"

func testErr(id int) (int, error) {

	if id < 0 {
		return 0, fmt.Errorf("id is illegal, id =%d", id)
	}
	return id, nil
}

func main() {

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
}
