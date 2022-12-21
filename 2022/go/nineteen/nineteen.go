package nineteen

import (
	"2022/go/utils"
	"fmt"
	"strings"
)

const input = "../input/19.txt"
const maxPart1 = 24
const maxPart2 = 32

func ParseInput() []Blueprint {
	lines := utils.ReadLines(input)
	bps := make([]Blueprint, len(lines))

	for i, l := range lines {
		l = strings.ReplaceAll(l, "Blueprint ", "")
		l = strings.ReplaceAll(l, ": Each ore robot costs ", ",")
		l = strings.ReplaceAll(l, " ore. Each clay robot costs ", ",")
		l = strings.ReplaceAll(l, " ore. Each obsidian robot costs ", ",")
		l = strings.ReplaceAll(l, " ore and ", ",")
		l = strings.ReplaceAll(l, " clay. Each geode robot costs ", ",")
		l = strings.ReplaceAll(l, " ore and ", ",")
		l = strings.ReplaceAll(l, " obsidian.", "")
		split := utils.Intify(utils.SplitStringOn(l, ","))

		max := Types2{0, 0, 0}

		if split[1] > max.or {
			max.or = split[1]
		}
		if split[2] > max.or {
			max.or = split[2]
		}
		if split[3] > max.or {
			max.or = split[3]
		}
		if split[4] > max.cl {
			max.cl = split[4]
		}
		if split[5] > max.or {
			max.or = split[5]
		}
		if split[6] > max.ob {
			max.ob = split[6]
		}

		bps[i] = Blueprint{
			split[0],
			Types2{split[1], 0, 0},
			Types2{split[2], 0, 0},
			Types2{split[3], split[4], 0},
			Types2{split[5], 0, split[6]},
			max,
		}
	}

	return bps
}

type Types struct {
	or int // ore
	cl int // clay
	ob int // obsidian
	ge int // geode
}

type Types2 struct {
	or int // ore
	cl int // clay
	ob int // obsidian
}

type Blueprint struct {
	id  int
	or  Types2 // ore
	cl  Types2 // clay
	ob  Types2 // obsidian
	ge  Types2 // geode
	max Types2 // max resource needed for a robot to be built
}

type State struct {
	m  int
	ro Types // robots
	re Types // resources
}

type Key struct {
	m  int
	ro Types
	re Types2
}

func (s State) copy() State {
	ro := Types{s.ro.or, s.ro.cl, s.ro.ob, s.ro.ge}
	re := Types{s.re.or, s.re.cl, s.re.ob, s.re.ge}
	return State{s.m, ro, re}
}

func (s State) toKey() Key {
	ro := Types{s.ro.or, s.ro.cl, s.ro.ob, s.ro.ge}
	re := Types2{s.re.or, s.re.cl, s.re.ob}
	return Key{s.m, ro, re}
}

func (s *State) produce() {
	s.m += 1
	s.re.or += s.ro.or
	s.re.cl += s.ro.cl
	s.re.ob += s.ro.ob
	s.re.ge += s.ro.ge
}

func (s State) oneGeodeRobotPerMinute(maxMinutes int) int {
	geodes := s.re.ge
	geodeRobots := s.ro.ge
	for m := s.m; m <= maxMinutes; m++ {
		geodes += geodeRobots
		geodeRobots += 1
	}
	return geodes
}

func (bp Blueprint) GetMax(minutes int) int {
	// BFS
	queue := []State{{0, Types{1, 0, 0, 0}, Types{0, 0, 0, 0}}}
	cache := map[Key]int{}
	max := 0
	for {
		if len(queue) == 0 {
			break
		}

		// pop
		st := queue[0]
		queue = queue[1:]
		k := st.toKey()

		// update cache
		geodes, seen := cache[k]

		if seen && geodes >= st.re.ge {
			continue
		}
		cache[k] = st.re.ge

		// Limit reached
		if st.m == minutes {
			if st.re.ge > max {
				max = st.re.ge
			}
			continue
		}

		// With optimistic nbr geode robots
		geodes = st.oneGeodeRobotPerMinute(minutes)

		// If geode robot can be produced every minute, optimistic is true :raised_hands:
		if st.ro.or == bp.ge.or && st.ro.ob == bp.ge.ob {
			if geodes > max {
				max = geodes
			}
		}

		// If optimistic is worse than current max, abandon
		if geodes < max {
			continue
		}

		// ore robot
		if st.re.or >= bp.or.or && st.ro.or <= bp.max.or {
			new := st.copy()
			new.produce()
			new.re.or -= bp.or.or
			new.ro.or += 1
			queue = append(queue, new)
		}

		// clay robot
		if st.re.or >= bp.cl.or && st.ro.cl <= bp.max.cl {
			new := st.copy()
			new.produce()
			new.re.or -= bp.cl.or
			new.ro.cl += 1
			queue = append(queue, new)
		}

		// obsidian robot
		if st.re.or >= bp.ob.or && st.re.cl >= bp.ob.cl && st.ro.ob <= bp.max.ob {
			new := st.copy()
			new.produce()
			new.re.or -= bp.ob.or
			new.re.cl -= bp.ob.cl
			new.ro.ob += 1
			queue = append(queue, new)
		}

		// geode robot
		if st.re.or >= bp.ge.or && st.re.ob >= bp.ge.ob {
			new := st.copy()
			new.produce()
			new.re.or -= bp.ge.or
			new.re.ob -= bp.ge.ob
			new.ro.ge += 1
			queue = append(queue, new)
		}

		// missing resource, but have robot producing it.
		oreMissing := st.re.or < bp.or.or || st.re.or < bp.cl.or || st.re.or < bp.ob.or || st.re.or < bp.ge.or
		clayMissing := st.re.cl < bp.ob.cl && st.ro.cl > 0
		obsidianMissing := st.re.ob < bp.ge.ob && st.ro.ob > 0

		// Build no robots
		if oreMissing || clayMissing || obsidianMissing {
			new := st.copy()
			new.produce()
			queue = append(queue, new)
		}

	}
	return max
}

func QualityLevels(bps []Blueprint) {
	sum := 0
	for _, bp := range bps {
		max := bp.GetMax(maxPart1)
		sum += bp.id * max
	}
	fmt.Println("Quality Level: ", sum)
}

func MaximumGeodes(bps []Blueprint) {
	prod := 1
	for _, bp := range bps {
		max := bp.GetMax(maxPart2)
		prod *= max
	}
	fmt.Println("Maximum product: ", prod)
}

func Run() {
	bps := ParseInput()
	QualityLevels(bps)
	MaximumGeodes(bps[:3])
}
