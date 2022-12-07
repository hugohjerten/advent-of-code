package main

import (
	"2022/go/five"
	"2022/go/four"
	"2022/go/one"
	"2022/go/seven"
	"2022/go/six"
	"2022/go/three"
	"2022/go/two"
	"log"
	"os"
)

var (
	solve = map[string]func(){
		"1": one.Run,
		"2": two.Run,
		"3": three.Run,
		"4": four.Run,
		"5": five.Run,
		"6": six.Run,
		"7": seven.Run,
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
