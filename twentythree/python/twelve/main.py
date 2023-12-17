"""Twelve."""
from functools import cache
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "12.txt"


Row = tuple[str, tuple]


def parse_input() -> list[Row]:
    """Parse input."""
    return [
        (line.split(" ")[0], tuple(int(g) for g in line.split(" ")[1].split(",")))
        for line in read_file(path)
    ]


def _find_combinations(condition: str) -> list[str]:
    """Find all possible combinations."""
    if condition.count("?") == 1:
        return [condition.replace("?", "."), condition.replace("?", "#")]

    combinations = _find_combinations(condition.replace("?", ".", 1))
    combinations.extend(_find_combinations(condition.replace("?", "#", 1)))

    return combinations


def _find_arrangements(row: Row) -> int:
    """Return number of possible arrangements."""
    combinations = _find_combinations(row[0])
    cnt = 0
    for c in combinations:
        group_sizes = tuple(len(g) for g in c.split(".") if g)

        if group_sizes == row[1]:
            cnt += 1

    return cnt


def part_1():
    """Part 1."""
    rows = parse_input()
    total = 0
    for r in rows:
        total += _find_arrangements(r)
    print("Part 1: ", total)


@cache
def _recurse(springs: str, clue: tuple, size: int = 0) -> int:  # noqa: PLR0911
    """Return number of possible arrangements."""
    if not springs:
        if len(clue) == 1 and size == clue[0]:
            return 1
        if len(clue) == 0 and size == 0:
            return 1

        return 0

    spring = springs[0]
    springs = springs[1:]

    if spring == "?":
        return _recurse("." + springs, clue, size) + _recurse("#" + springs, clue, size)

    if spring == "#":
        # If size bigger than allowed clue, return 0
        if len(clue) == 0 or size > clue[0]:
            return 0

        # Inside group, increment size.
        return _recurse(springs, clue, size + 1)

    if spring == ".":
        # If group not started (0) and size not equal to clue, return 0
        if size != 0 and (len(clue) == 0 or size != clue[0]):
            return 0

        # Start new size. If previous size never started, pass all clues. Else one less.
        return _recurse(springs, clue[1:] if size != 0 else clue, 0)

    return 0


def part_2():
    """Part 2."""
    rows = [("?".join([row[0]] * 5), row[1] * 5) for row in parse_input()]
    total = 0
    for i, r in enumerate(rows):
        total += _recurse(r[0], r[1])
    print("Part 2: ", total)


def main():
    """Main."""
    part_1()
    part_2()
