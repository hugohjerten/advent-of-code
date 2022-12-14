package six

import (
	"2022/go/utils"
	"fmt"
)

type Type int

const (
	Start   Type = 4
	Message Type = 14
)

type SubRoutine struct {
	cache   []byte
	signal  string
	current int
	t       Type
}

func NewSubRoutine(signal string, t Type) SubRoutine {
	cache := make([]byte, int(t))

	for i := 0; i < int(t); i++ {
		cache[i] = signal[i]
	}
	return SubRoutine{cache, signal, int(t) - 1, t}
}

func (s SubRoutine) packetStart() bool {
	// Create a set out of the cache
	set := map[byte]struct{}{}
	for _, b := range s.cache {
		set[b] = struct{}{}
	}

	return len(set) == int(s.t)
}

func (s *SubRoutine) next() {
	s.current = s.current + 1
	s.cache = s.cache[1:]
	s.cache = append(s.cache, s.signal[s.current])
}

func (s SubRoutine) Find() int {
	for {
		if s.current > len(s.signal) {
			break
		}
		if s.packetStart() {
			return s.current + 1
		}
		s.next()
	}
	panic("Did not find it.")
}

const input = "../input/6.txt"

func Run() {
	signal := utils.ReadLines(input)[0]
	s := NewSubRoutine(signal, Start)
	fmt.Println("Number of characters to process: ", s.Find())

	s = NewSubRoutine(signal, Message)
	fmt.Println("Number of characters to process: ", s.Find())
}
