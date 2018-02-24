package main

import "fmt"

func main() {

	var str string = "HelloWorld"
	var str2 string = "你好，世界"

	fmt.Println("len of str =", len(str), "len of str2 =", len(str2))
	fmt.Println("len of str splice =", len([]rune(str)), "len of str2 splice =", len([]rune(str2)))
	fmt.Println("str[1:2] =", str[1:2], "str2[1:2] =", string([]rune(str2)[1]))
}
