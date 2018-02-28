package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func add(result *int, group *sync.WaitGroup) {
	for i := 0; i <= 10000; i++ {
		mutex.Lock()
		*result += i
		mutex.Unlock()
	}
	(*group).Done()
}

func main() {

	const counter = 2
	result := 0
	var group sync.WaitGroup

	group.Add(counter)

	for i := 0; i < counter; i++ {
		go add(&result, &group)
	}

	group.Wait()
	fmt.Println(result)
}
