"""Fifteen."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "15.txt"


def parse_input() -> list[str]:
    """Parse input."""
    return read_file(path)[0].split(",")


def hash_algorithm(step: str) -> int:
    """Run hash algorithm on step."""
    val = 0
    for c in step:
        val += ord(c)
        val *= 17
        val = val % 256

    return val


def part_1():
    """Part 1."""
    sequence = parse_input()
    total = sum([hash_algorithm(c) for c in sequence])

    print("Part 1: ", total)


def main():
    """Main."""
    part_1()
