"""Twelve."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "12.txt"


Row = tuple[str, list[int]]


def parse_input() -> list[Row]:
    """Parse input."""
    return [
        (line.split(" ")[0], [int(g) for g in line.split(" ")[1].split(",")])
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
        group_sizes = [len(g) for g in c.split(".") if g]

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


def main():
    """Main."""
    part_1()
