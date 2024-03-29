package main

import (
	"log"
	"os"
	"twentythree/go/one"
	"twentythree/go/three"
	"twentythree/go/two"
)

var (
	solve = map[string]func(){
		"1": one.Run,
		"2": two.Run,
		"3": three.Run,
	}
)

func main() {
	if len(os.Args) != 2 {
		panic("Provide exactly one day!")
	}

	day := os.Args[1]
	if _, ok := solve[day]; !ok {
		log.Fatal("Bad day input: ", day)
		panic("FAILURE")
	}

	solve[day]()
}
