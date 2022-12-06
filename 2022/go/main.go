package main

import (
	"2022/go/five"
	"2022/go/four"
	"2022/go/one"
	"2022/go/six"
	"2022/go/three"
	"2022/go/two"
	"log"
	"os"
)

var (
	solve = map[string]func(){
		"one":   one.Run,
		"two":   two.Run,
		"three": three.Run,
		"four":  four.Run,
		"five":  five.Run,
		"six":   six.Run,
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
