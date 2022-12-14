package thirteen

import (
	"2022/go/utils"
	"fmt"
	"sort"
	"strconv"
)

const input = "../input/13.txt"

type Integer struct {
	val int
}

type Packet struct {
	val string
	idx int
}

type Comparison int

const (
	BothInt Comparison = iota
	BothList
	BothEmpty
	LeftEmpty
	RightEmpty
	LeftList
	RightList
	BothNil
)

type Order int

const (
	Right Order = iota
	Wrong
	Same
)

func (p Packet) nestedList(start int) string {
	nstd := 0
	for i := start; i < len(p.val); i++ {
		switch p.val[i] {
		case 91: // '['
			nstd += 1
		case 93:
			nstd -= 1

			if nstd == 0 {
				sub := p.val[start : i+1]
				return sub
			}
		}
	}
	panic("Couldn't get slice.")
}

func (p *Packet) Next() (*Integer, *Packet) {
	if p.idx >= len(p.val) {
		return nil, nil
	}

	r := p.val[p.idx]

	switch r {
	case 91: // '['
		val := p.nestedList(p.idx)

		p.idx += len(val)
		return nil, &Packet{val, 1}

	case 93: // ']'
		p.idx += 1
		return p.Next()

	case 44: // ','
		p.idx += 1
		return p.Next()

	default:
		var val int
		next := p.val[p.idx+1]

		if next >= 48 && next <= 57 {
			// i.e. if value is two digits, next digit is between 0<9
			val, _ = strconv.Atoi(p.val[p.idx : p.idx+2])
			p.idx += 2

		} else {
			// i.e. value is one digit long
			val, _ = strconv.Atoi(string(r))
			p.idx += 1
		}

		return &Integer{val}, nil
	}

}

type Pair struct {
	l Packet
	r Packet
}

func ToPacket(i int) Packet {
	return Packet{("[" + strconv.Itoa(i) + "]"), 1}
}

func (p Packet) IsEmpty() bool {
	return p.val == "[]"
}

func determineComparison(li *Integer, ri *Integer, lp *Packet, rp *Packet) Comparison {
	if li != nil && ri != nil {
		return BothInt
	}

	if lp != nil && rp != nil {
		if lp.IsEmpty() && rp.IsEmpty() {
			return BothEmpty
		}
		if lp.IsEmpty() {
			return LeftEmpty
		}
		if rp.IsEmpty() {
			return RightEmpty
		}
		return BothList
	}

	if li == nil && lp == nil && ri == nil && rp == nil {
		return BothNil
	}

	if li == nil && lp == nil {
		return LeftEmpty
	}

	if ri == nil && rp == nil {
		return RightEmpty
	}

	if lp != nil && ri != nil {
		return LeftList
	}

	if li != nil && rp != nil {
		return RightList
	}

	panic("Can't determine comparison.")
}

func (p *Pair) compare() Order {
	for {
		li, lp := p.l.Next()
		ri, rp := p.r.Next()

		switch determineComparison(li, ri, lp, rp) {
		case BothInt:
			if li.val > ri.val {
				return Wrong
			} else if li.val < ri.val {
				return Right
			}
		case BothList:
			nested := Pair{*lp, *rp}
			order := nested.compare()
			if order == Right || order == Wrong {
				return order
			}

		case BothNil:
			return Same
		case LeftEmpty:
			return Right
		case RightEmpty:
			return Wrong

		case LeftList:
			cnvrtd := ToPacket(ri.val)
			nested := Pair{*lp, cnvrtd}
			order := nested.compare()
			if order == Right || order == Wrong {
				return order
			}

		case RightList:
			cnvrtd := ToPacket(li.val)
			nested := Pair{cnvrtd, *rp}
			order := nested.compare()
			if order == Right || order == Wrong {
				return order
			}
		}
	}
}

func ComparePairs(pairs []Pair) {
	sum := 0
	for i, p := range pairs {
		if p.compare() == Right {
			sum += i + 1
		}
	}

	fmt.Println("Sum of indices: ", sum)
}

func ParsePairs(input string) []Pair {
	split := utils.SeparateSliceOnNewLine(utils.ReadLines(input))
	pairs := make([]Pair, len(split))
	for i, p := range split {
		pairs[i] = Pair{Packet{p[0], 1}, Packet{p[1], 1}}
	}

	return pairs
}

func ParsePackets(input string) []Packet {
	lines := utils.RemoveEmptyLines(utils.ReadLines(input))
	ps := make([]Packet, len(lines))
	for i, l := range lines {
		ps[i] = Packet{l, 1}
	}

	ps = append(ps, []Packet{{"[[2]]", 1}, {"[[6]]", 1}}...)
	return ps
}

func DecoderKey(ps []Packet) {
	sort.Slice(ps, func(i, j int) bool {
		pair := Pair{ps[i], ps[j]}
		return pair.compare() == Right
	})

	prod := 1
	for i, p := range ps {
		if p.val == "[[2]]" || p.val == "[[6]]" {
			prod *= (i + 1)
		}
	}
	fmt.Println("DecoderKey: ", prod)
}

func Run() {
	pairs := ParsePairs(input)
	ComparePairs(pairs)

	packets := ParsePackets(input)
	DecoderKey(packets)

}
