package main

import (
	"2022/go/eight"
	"2022/go/eighteen"
	"2022/go/eleven"
	"2022/go/fifteen"
	"2022/go/five"
	"2022/go/four"
	"2022/go/fourteen"
	"2022/go/nine"
	"2022/go/nineteen"
	"2022/go/one"
	"2022/go/seven"
	"2022/go/seventeen"
	"2022/go/six"
	"2022/go/sixteen"
	"2022/go/ten"
	"2022/go/thirteen"
	"2022/go/three"
	"2022/go/twelve"
	"2022/go/twenty"
	"2022/go/twentyone"
	"2022/go/twentytwo"
	"2022/go/two"
	"log"
	"os"
)

var (
	solve = map[string]func(){
		"1":  one.Run,
		"2":  two.Run,
		"3":  three.Run,
		"4":  four.Run,
		"5":  five.Run,
		"6":  six.Run,
		"7":  seven.Run,
		"8":  eight.Run,
		"9":  nine.Run,
		"10": ten.Run,
		"11": eleven.Run,
		"12": twelve.Run,
		"13": thirteen.Run,
		"14": fourteen.Run,
		"15": fifteen.Run,
		"16": sixteen.Run,
		"17": seventeen.Run,
		"18": eighteen.Run,
		"19": nineteen.Run,
		"20": twenty.Run,
		"21": twentyone.Run,
		"22": twentytwo.Run,
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
