"""Eleven."""
import itertools
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "11.txt"


Image = list[str]
Coordinates = tuple[int, int]
Pair = tuple[Coordinates, Coordinates]


def parse_input() -> list[Coordinates]:
    """Parse input."""
    image = read_file(path)
    rows = [r for r, row in enumerate(image) if "#" not in row]
    cols = []
    for c in range(len(image[0])):
        if all(image[r][c] == "." for r in range(len(image))):
            cols.append(c)

    # Add expansions
    for r in rows[::-1]:
        image.insert(r, "." * len(image[0]))

    for c in cols[::-1]:
        for r in range(len(image)):
            image[r] = image[r][:c] + "." + image[r][c:]

    # Return galaxies
    return [
        Coordinates((y, x))
        for y in range(len(image))
        for x in range(len(image[0]))
        if image[y][x] == "#"
    ]


def part_1():
    """Part 1."""
    galaxies = parse_input()

    # Manhatten distances
    distances = [
        abs(g1[0] - g2[0]) + abs(g1[1] - g2[1])
        for g1, g2 in itertools.combinations(galaxies, 2)
    ]

    print("Part 1: ", sum(distances))


def main():
    """Main."""
    part_1()
