"""Fourteen."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "14.txt"

# row, col
Coordinate = tuple[int, int]

CYCLES = 1_000_000_000


def parse_input() -> list[list[str]]:
    """Parse input."""
    return [[s for s in line] for line in read_file(path)]


def _slide_rock_north(r: int, c: int, platform: list[str]) -> list[str]:
    """Slide rock north."""
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


def _slide_rock_south(r: int, c: int, platform: list[str]) -> list[str]:
    """Slide rock south."""
    while True:
        below = r + 1
        if below == len(platform):
            return platform

        if platform[below][c] == ".":
            platform[below][c] = "O"
            platform[r][c] = "."
        else:
            return platform

        r = below


def _slide_rock_east(r: int, c: int, platform: list[str]) -> list[str]:
    """Slide rock east."""
    while True:
        right = c + 1
        if right == len(platform[0]):
            return platform

        if platform[r][right] == ".":
            platform[r][right] = "O"
            platform[r][c] = "."
        else:
            return platform

        c = right


def _slide_rock_west(r: int, c: int, platform: list[str]) -> list[str]:
    """Slide rock west."""
    while True:
        left = c - 1
        if left < 0:
            return platform

        if platform[r][left] == ".":
            platform[r][left] = "O"
            platform[r][c] = "."
        else:
            return platform

        c = left


def _tilt_north(platform: list[str]) -> list[str]:
    """Tilt north."""
    for r in range(1, len(platform)):
        for c in range(len(platform[0])):
            if platform[r][c] == "O":
                platform = _slide_rock_north(r, c, platform)
    return platform


def _tilt_west(platform: list[str]) -> list[str]:
    """Tilt west."""
    for c in range(1, len(platform[0])):
        for r in range(len(platform)):
            if platform[r][c] == "O":
                platform = _slide_rock_west(r, c, platform)
    return platform


def _tilt_south(platform: list[str]) -> list[str]:
    """Tilt south."""
    for r in range(len(platform) - 2, -1, -1):
        for c in range(len(platform[0])):
            if platform[r][c] == "O":
                platform = _slide_rock_south(r, c, platform)
    return platform


def _tilt_east(platform: list[str]) -> list[str]:
    """Tilt east."""
    for c in range(len(platform[0]) - 2, -1, -1):
        for r in range(len(platform)):
            if platform[r][c] == "O":
                platform = _slide_rock_east(r, c, platform)
    return platform


def part_1():
    """Part 1."""
    platform = _tilt_north(parse_input())

    # Calculate load
    loads = [
        x * (i + 1) for i, x in enumerate([line.count("O") for line in platform][::-1])
    ]
    print("Part 1:", sum(loads))


def _rock_locations(platform: list[str]) -> tuple[Coordinate]:
    """Get a tuple of all rock locations."""
    rocks: list[Coordinate] = []
    for r in range(len(platform)):
        for c in range(len(platform[0])):
            if platform[r][c] == "O":
                rocks.append(Coordinate((r, c)))

    return tuple(rocks)


def part_2():
    """Part 2."""
    platform = parse_input()
    states: dict[tuple, int] = {}

    i = 0
    cycle_found = False
    while i < CYCLES:
        platform = _tilt_north(platform)
        platform = _tilt_west(platform)
        platform = _tilt_south(platform)
        platform = _tilt_east(platform)

        state = _rock_locations(platform)

        if state in states and not cycle_found:
            distance_to_goal = CYCLES - i
            loop_length = i - states[state]
            print(f"Distance to goal: {distance_to_goal}, loop_length: {loop_length}")
            i = CYCLES - distance_to_goal % loop_length
            cycle_found = True

        states[state] = i
        i += 1

    # Calculate load
    loads = [
        x * (i + 1) for i, x in enumerate([line.count("O") for line in platform][::-1])
    ]
    print("Part 2:", sum(loads))


def main():
    """Main."""
    part_1()
    part_2()
