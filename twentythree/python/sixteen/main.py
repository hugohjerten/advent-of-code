"""Sixteen."""
from pathlib import Path
from typing import NamedTuple

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "16.txt"


Grid = list[list[str]]
Coordinate = tuple[int, int]


class Beam(NamedTuple):
    """Beam."""

    coordinate: Coordinate
    direction: str  # i.e. one of (N, E, S, W)


def parse_input() -> Grid:
    """Parse input."""
    return [[c for c in line] for line in read_file(path)]


def _next(current: Coordinate, direction: str) -> Beam:
    """Find next beam location(s)."""
    if direction == "N":
        return Beam((current[0] - 1, current[1]), direction)
    if direction == "E":
        return Beam((current[0], current[1] + 1), direction)
    if direction == "S":
        return Beam((current[0] + 1, current[1]), direction)
    return Beam((current[0], current[1] - 1), direction)


DIRECTION = {
    "/": {"N": "E", "E": "N", "S": "W", "W": "S"},
    "\\": {"N": "W", "E": "S", "S": "E", "W": "N"},
    "|": {"N": "N", "S": "S"},
    "-": {"E": "E", "W": "W"},
}


def traverse(grid: Grid) -> Grid:
    """Traverse."""
    nexts: list[Beam] = [Beam((0, 0), "E")]
    energised = [[0 for _ in range(len(grid[0]))] for _ in range(len(grid))]
    processed: set[Beam] = set()

    while nexts:
        beam = nexts.pop(0)
        if beam in processed:
            continue

        processed.add(beam)
        crd = beam.coordinate
        if crd[0] < 0 or crd[0] == len(grid) or crd[1] < 0 or crd[1] == len(grid[0]):
            continue

        energised[crd[0]][crd[1]] = 1
        space = grid[crd[0]][crd[1]]

        if space == ".":
            nexts.append(_next(crd, beam.direction))

        elif space == "|" and beam.direction in ("E", "W"):
            nexts.append(_next(crd, "N"))
            nexts.append(_next(crd, "S"))

        elif space == "-" and beam.direction in ("N", "S"):
            nexts.append(_next(crd, "E"))
            nexts.append(_next(crd, "W"))

        else:
            nexts.append(_next(crd, DIRECTION[space][beam.direction]))

    return energised


def part_1():
    """Part 1."""
    grid = parse_input()
    energised = traverse(grid)

    total = sum([sum(row) for row in energised])
    print("Part 1: ", total)


def main():
    """Main."""
    part_1()
