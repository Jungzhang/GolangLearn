package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "误报,偶发性问题,历史问题,问题影响可接受带问题上线"
	ss := strings.Split(str, ",")
	for _, s := range ss{
		fmt.Println(s)
	}
}
