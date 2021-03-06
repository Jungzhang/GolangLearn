package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var mmm = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var mm = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mm)

	for v := range mm {
		fmt.Println("key =", v, "   value =", mm[v])
	}
}
