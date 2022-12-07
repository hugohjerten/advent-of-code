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
		if lines[i] == "$ ls" {
			i = d.parseLsOutput(lines, i+1)
		}

		if i == len(lines) {
			break
		}

		if lines[i] == "$ cd .." {
			d = d.parent

			i++
			continue
		}

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

func (d Dir) Size() int {
	sum := 0
	for _, size := range d.files {
		sum += size
	}

	for _, dir := range d.children {
		sum += dir.Size()
	}

	return sum
}

func (d Dir) Sizes() []int {
	sizes := make([]int, 0)
	sizes = append(sizes, d.Size())

	for _, child := range d.children {
		sizes = append(sizes, child.Sizes()...)
	}

	return sizes
}

func (d Dir) SumOfSizes() {
	sizes := d.Sizes()
	total := 0

	for _, size := range sizes {
		if size < 100000 {
			total += size
		}
	}

	fmt.Println("Sum of sizes: ", total)
}

func (d Dir) SmallestDir() {
	need := 30000000 - (70000000 - d.Size())
	dirs := d.Sizes()

	size := 10000000
	for _, s := range dirs {
		if s >= need && s < size {
			size = s
		}
	}

	fmt.Println("Smallest dir: ", size)
}

const input = "../input/7.txt"

func Run() {
	lines := utils.ReadLines(input)
	root := readFileSystem(lines)

	root.SumOfSizes()
	root.SmallestDir()
}
