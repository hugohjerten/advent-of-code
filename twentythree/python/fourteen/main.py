"""Fourteen."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "14.txt"


def parse_input() -> list[list[str]]:
    """Parse input."""
    return [[s for s in line] for line in read_file(path)]


def _slide_rock(r: int, c: int, platform: list[str]) -> list[str]:
    """Slide rock."""
    while True:
        above = r - 1
        if above < 0:
            return platform

        if platform[above][c] == ".":
            platform[above][c] = "O"
            platform[r][c] = "."
        else:
            return platform

        r = above


def tilt_north(platform: list[str]) -> list[str]:
    """Tilt platform north."""
    for r in range(1, len(platform)):
        for c in range(len(platform[0])):
            if platform[r][c] == "O":
                platform = _slide_rock(r, c, platform)

    return platform


def part_1():
    """Part 1."""
    platform = parse_input()
    tilted = tilt_north(platform)
    loads = [
        x * (i + 1) for i, x in enumerate([line.count("O") for line in tilted][::-1])
    ]
    print("Part 1:", sum(loads))


def main():
    """Main."""
    part_1()
