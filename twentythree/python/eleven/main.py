"""Eleven."""
import itertools
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "11.txt"


Image = list[str]
Coordinates = tuple[int, int]
Pair = tuple[Coordinates, Coordinates]


def parse_input_part_1() -> list[Coordinates]:
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
    galaxies = parse_input_part_1()

    # Manhatten distances
    distances = [
        abs(g1[0] - g2[0]) + abs(g1[1] - g2[1])
        for g1, g2 in itertools.combinations(galaxies, 2)
    ]

    print("Part 1: ", sum(distances))


def parse_input_part_2() -> tuple[list[Coordinates], list[int], list[int]]:
    """Parse input."""
    image = read_file(path)
    empty_rows = [r for r, row in enumerate(image) if "#" not in row]
    empty_cols = []
    for c in range(len(image[0])):
        if all(image[r][c] == "." for r in range(len(image))):
            empty_cols.append(c)

    galaxies = [
        Coordinates((y, x))
        for y in range(len(image))
        for x in range(len(image[0]))
        if image[y][x] == "#"
    ]

    return galaxies, empty_rows, empty_cols


def part_2():
    """Part 2."""
    galaxies, empty_rows, empty_cols = parse_input_part_2()

    # Manhatten distances
    distances: list[int] = []
    for g1, g2 in itertools.combinations(galaxies, 2):
        min_y, max_y = min(g1[0], g2[0]), max(g1[0], g2[0])
        min_x, max_x = min(g1[1], g2[1]), max(g1[1], g2[1])

        # Add extra distance for each empty row and column
        extra_rows = len([r for r in empty_rows if min_y < r < max_y]) * (1000000 - 1)
        extra_cols = len([c for c in empty_cols if min_x < c < max_x]) * (1000000 - 1)

        distances.append(
            abs(g1[0] - g2[0]) + extra_rows + abs(g1[1] - g2[1]) + extra_cols
        )

    print("Part 2: ", sum(distances))


def main():
    """Main."""
    part_1()
    part_2()
