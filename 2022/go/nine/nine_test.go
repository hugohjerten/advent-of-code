package nine_test

import (
	"2022/go/nine"
	"2022/go/utils"
	"testing"
)

func Test2Knots(t *testing.T) {

	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	r := nine.NewRope(2)
	ds := nine.ParseDirections(utils.SplitStringOnNewline(input))
	got := r.GetPositions(ds)
	want := 13

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test10Knots(t *testing.T) {

	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	r := nine.NewRope(10)
	ds := nine.ParseDirections(utils.SplitStringOnNewline(input))
	got := r.GetPositions(ds)
	want := 36

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
