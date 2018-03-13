package main

import (
	"flag"
	"fmt"
)

func main() {

	score := flag.Int("s", 60, "请输入分数")
	flag.Parse()
	fmt.Println("score :", *score)

}
