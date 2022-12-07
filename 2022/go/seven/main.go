package seven

import (
	"2022/go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Dir struct {
	name     string
	parent   *Dir
	children map[string]*Dir
	files    []int
}

func NewDir(n string, p *Dir) *Dir {
	return &Dir{n, p, make(map[string]*Dir), make([]int, 0)}
}

func (d *Dir) parseLsOutput(lines []string, i int) int {
	for {
		if i == len(lines) {
			break
		}
		split := utils.SplitStringOnWhitespace(lines[i])

		// ls output is done
		if split[0] == "$" {
			break
		}

		// add directory
		if split[0] == "dir" {
			d.children[split[1]] = NewDir(split[1], d)

			i++
			continue
		}

		// add file
		size, _ := strconv.Atoi(split[0])
		d.files = append(d.files, size)

		i++
	}

	return i
}

func readFileSystem(lines []string) *Dir {
	root := NewDir("/", nil)
	d := root

	i := 1
	for {
		// ls
		if lines[i] == "$ ls" {
			i = d.parseLsOutput(lines, i+1)
		}

		if i == len(lines) {
			break
		}

		// cd ..
		if lines[i] == "$ cd .." {
			d = d.parent

			i++
			continue
		}

		// cd subdir
		if strings.HasPrefix(lines[i], "$ cd ") {
			split := utils.SplitStringOnWhitespace(lines[i])
			d = d.children[split[2]]

			i++
			continue
		}

		i++
	}

	return root
}

func (d Dir) SumOfSizes(total int) (int, int) {
	sum := 0
	for _, size := range d.files {
		sum += size
	}

	childSum := 0
	for _, dir := range d.children {
		childSum, total = dir.SumOfSizes(total)
		sum += childSum
	}

	if sum <= 100000 {
		total += sum
	}

	return sum, total
}

const input = "../input/7.txt"

func Run() {
	lines := utils.ReadLines(input)
	root := readFileSystem(lines)

	_, total := root.SumOfSizes(0)
	fmt.Println("Sum of sizes: ", total)
}
