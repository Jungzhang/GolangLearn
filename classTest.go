package main

import "fmt"

type person struct {
	name string
}

type employee struct {
	id int
	person
}

func (p *person)modify (name string)  {
	p.name = name
}

func (e *employee) modify(id int, name string) {
	e.id = id
	e.person.modify(name)
}


func main() {

	e := employee{id:1}
	fmt.Println(e)
	e.modify(2, "Raiden")
	fmt.Println(e)
}
