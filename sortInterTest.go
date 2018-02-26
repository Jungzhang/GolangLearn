package main

import (
	"sort"
	"fmt"
)

type StringSplice []string

func (s StringSplice) Len() int {
	return len(s)
}

func (s StringSplice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSplice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	s := StringSplice{"zhanggen", "Jung", "Raiden"}
	fmt.Println(s)

	sort.Sort(s)
	fmt.Println(s)

}
