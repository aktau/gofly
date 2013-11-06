package main

import (
	"fmt"
	. "github.com/aktau/gofly/flightengine"
	"log"
	"time"
)

func main() {
	fmt.Printf("gofly v%d.%d\n", 0, 1)

	m := Momondo{}
	dep := time.Date(2014, time.November, 10, 23, 0, 0, 0, time.UTC)
	ret := time.Date(2014, time.November, 10, 23, 0, 0, 0, time.UTC)
	priceSpec, err := m.Price("BRU", "LIM", dep, ret, "true")
	if err != nil {
		log.Fatal(err)
	}

	price03, err := m.PriceAvg("BRU", "LIM", 3, 2014)
	if err != nil {
		log.Fatal(err)
	}

	price04, err := m.PriceAvg("BRU", "LIM", 4, 2014)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("specific price for 03/2014 => %f\n", priceSpec)
	fmt.Printf("average price for 03/2014 => %f\n", price03)
	fmt.Printf("average price for 04/2014 => %f\n", price04)
}
