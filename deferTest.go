package main

import (
	"fmt"
	"time"
	"log"
)

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		//r = t + 5
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}


func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the
	time.Sleep(10 * time.Second)      // simulate slow
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	fmt.Println("f1 =", f1())
	fmt.Println("f2 =", f2())
	fmt.Println("f3 =", f3())
	bigSlowOperation()
}
