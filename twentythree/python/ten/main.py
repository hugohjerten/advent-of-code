"""Ten."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "10.txt"


Coordinate = tuple[int, int]


class Grid:
    """Grid."""

    grid: list[str]
    start: Coordinate

    def __init__(self) -> None:
        """Grid."""
        self.grid = read_file(path)
        for row, line in enumerate(self.grid):
            if "S" in line:
                self.start = Coordinate((row, line.index("S")))
                return

        raise Exception("Shouldn't be here.")

    def _ok_coordinate(self, p: Coordinate, previous: Coordinate | None) -> bool:
        """Check if valid coordinate."""
        return not (
            p == previous
            or p[0] < 0
            or p[0] == len(self.grid)
            or p[1] < 0
            or p[1] == len(self.grid[0])
            or self.grid[p[0]][p[1]] == "."
        )

    def _find_next_coordinate(
        self, p: Coordinate, previous: Coordinate | None
    ) -> Coordinate:
        """Find next coordinate."""
        # Check north
        if self.grid[p[0]][p[1]] in "S|LJ":
            q = Coordinate((p[0] - 1, p[1]))
            if self._ok_coordinate(q, previous) and (self.grid[q[0]][q[1]] in "S|F7"):
                return q

        # Check east
        if self.grid[p[0]][p[1]] in "S-LF":
            q = Coordinate((p[0], p[1] + 1))
            if self._ok_coordinate(q, previous) and (self.grid[q[0]][q[1]] in "S-J7"):
                return q

        # Check south
        if self.grid[p[0]][p[1]] in "S|7F":
            q = Coordinate((p[0] + 1, p[1]))
            if self._ok_coordinate(q, previous) and (self.grid[q[0]][q[1]] in "S|JL"):
                return q

        # Check west
        if self.grid[p[0]][p[1]] in "S-J7":
            q = Coordinate((p[0], p[1] - 1))
            if self._ok_coordinate(q, previous) and (self.grid[q[0]][q[1]] in "S-FL"):
                return q

        raise Exception("Shouldn't be here.")

    def traverse(self) -> int:
        """Traverse pipe."""
        previous = None
        current = self.start
        cnt = 0

        while True:
            next_ = self._find_next_coordinate(current, previous)
            cnt += 1

            if self.grid[next_[0]][next_[1]] == "S":
                break

            previous = current
            current = next_

        return int(cnt / 2)


def part_1():
    """Part 1."""
    grid = Grid()
    steps = grid.traverse()
    print("Part 1: ", steps)


def main():
    """Main."""
    part_1()
