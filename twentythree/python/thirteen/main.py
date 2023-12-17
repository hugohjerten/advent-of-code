"""Thirteen."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "13.txt"


def parse_input():
    """Parse input."""
    lines = read_file(path)
    ps: list[list[str]] = [[]]
    for line in lines:
        if not line:
            ps.append([])
            continue
        ps[-1].append(line)

    return ps


def _horizontal(r: int, p: list[str]) -> bool:
    """Check horizontal."""
    for k in range(r):
        if r + k == len(p):
            # Edge reached
            return True

        if p[r + k] != p[r - 1 - k]:
            return False

    return True


def _vertical(c: int, p: list[str]) -> bool:
    """Check vertical."""
    left, right = "", ""

    for k in range(c):
        if c + k == len(p[0]):
            # Edge reached
            return True

        for r in range(len(p)):
            left += p[r][c - 1 - k]
            right += p[r][c + k]

        if left != right:
            return False

        left, right = "", ""

    return True


def find_mirror(p: list[str]) -> int:
    """Find mirror."""
    for j in range(1, len(p)):
        if _horizontal(j, p):
            return j * 100

    for j in range(1, len(p[0])):
        if _vertical(j, p):
            return j

    raise Exception("Shouldn't be here.")


def part_1():
    """Part 1."""
    patterns = parse_input()
    total = 0
    for p in patterns:
        total += find_mirror(p)

    print("Part 1: ", total)


def main():
    """Main."""
    part_1()
