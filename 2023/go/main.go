package main

import (
	"2023/go/one"
	"2023/go/two"
	"log"
	"os"
)

var (
	solve = map[string]func(){
		"1": one.Run,
		"2": two.Run,
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
