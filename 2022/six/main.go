package six

import (
	"2022/utils"
	"fmt"
)

type SubRoutine struct {
	cache   []byte
	signal  string
	current int
}

func NewSubRoutine(signal string) SubRoutine {
	mrkr := 4
	cache := make([]byte, mrkr)

	for i := 0; i < mrkr; i++ {
		cache[i] = signal[i]
	}
	return SubRoutine{cache, signal, 3}
}

func (s SubRoutine) packetStart() bool {
	// Create a set out of the cache
	set := map[byte]struct{}{}
	for _, b := range s.cache {
		set[b] = struct{}{}
	}

	// If length of set is not 4, not all characters are different
	return len(set) == 4
}

func (s *SubRoutine) next() {
	s.current = s.current + 1
	s.cache = s.cache[1:]
	s.cache = append(s.cache, s.signal[s.current])
}

func (s SubRoutine) FindPacketStart() int {
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

func Run(filePath string) {
	signal := utils.ReadLines(filePath)[0]
	s := NewSubRoutine(signal)
	markerEnd := s.FindPacketStart()
	fmt.Println("Number of characters to process: ", markerEnd)

}
