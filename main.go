package main

import (
	"fmt"
	. "github.com/aktau/gofly/flightengine"
	"log"
)

func main() {
	fmt.Printf("gofly v%d.%d\n", 0, 1)

	m := Momondo{}
	price03, err := m.AvgPrice("BRU", "LIM", 3, 2014)
	if err != nil {
		log.Fatal(err)
	}

	price04, err := m.AvgPrice("BRU", "LIM", 4, 2014)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("average price for 03/2014 => %f\n", price03)
	fmt.Printf("average price for 04/2014 => %f\n", price04)
}
