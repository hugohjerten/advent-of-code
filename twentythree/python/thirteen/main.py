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


def _one_unequal(left: str, right: str) -> bool:
    """If exactly one unequal."""
    unequal_found = False
    for i in range(len(left)):
        if left[i] != right[i]:
            if unequal_found:
                return False
            unequal_found = True

    return unequal_found


def _horizontal(r: int, p: list[str], part_2: bool) -> bool:  # noqa: PLR0911
    """Check horizontal."""
    unequal_found = False
    for k in range(r):
        if r + k == len(p):
            # Edge reached
            if part_2 and not unequal_found:
                return False

            return True

        above = p[r - 1 - k]
        below = p[r + k]

        if part_2:
            if _one_unequal(below, above):
                if unequal_found:
                    return False
                unequal_found = True

            elif below != above:
                return False

        elif below != above:
            return False

    if part_2 and not unequal_found:
        return False

    return True


def _vertical(c: int, p: list[str], part_2: bool) -> bool:  # noqa: C901, PLR0911
    """Check vertical."""
    unequal_found = False
    left, right = "", ""

    for k in range(c):
        if c + k == len(p[0]):
            # Edge reached
            if part_2 and not unequal_found:
                return False

            return True

        for r in range(len(p)):
            left += p[r][c - 1 - k]
            right += p[r][c + k]

        if part_2:
            if _one_unequal(left, right):
                if unequal_found:
                    return False
                unequal_found = True

            elif left != right:
                return False

        elif left != right:
            return False

        left, right = "", ""

    if part_2 and not unequal_found:
        return False

    return True


def find_mirror(p: list[str], part_2: bool = False) -> int:
    """Find mirror."""
    for j in range(1, len(p)):
        if _horizontal(j, p, part_2):
            return j * 100

    for j in range(1, len(p[0])):
        if _vertical(j, p, part_2):
            return j

    raise Exception("Shouldn't be here.")


def part_1():
    """Part 1."""
    patterns = parse_input()
    total = 0
    for p in patterns:
        total += find_mirror(p)

    print("Part 1: ", total)


def part_2():
    """Part 2."""
    patterns = parse_input()
    total = 0
    for p in patterns:
        total += find_mirror(p, part_2=True)

    print("Part 2: ", total)


def main():
    """Main."""
    part_1()
    part_2()
