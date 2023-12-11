"""Six."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "6.txt"
Race = tuple[int, int]


def parse_input() -> list[Race]:
    """Parse input."""
    lines = [[int(k) for k in line.split(" ")[1:] if k] for line in read_file(path)]
    return [(lines[0][i], lines[1][i]) for i in range(len(lines[0]))]


def get_ways_to_win(race: Race) -> int:
    """Get number of ways to win race."""
    ways_to_win = 0
    for speed in range(race[0] + 1):
        ms_left = race[0] - speed
        if speed * ms_left > race[1]:
            ways_to_win += 1
    return ways_to_win


def part_1():
    """Part 1."""
    lines = [[int(k) for k in line.split(" ")[1:] if k] for line in read_file(path)]
    races = [(lines[0][i], lines[1][i]) for i in range(len(lines[0]))]
    product = 1
    for race in races:
        product *= get_ways_to_win(race)

    print("Part 1: ", product)


def part_2():
    """Part 2."""
    lines = read_file(path)
    race = (
        int(lines[0].replace(" ", "").split(":")[1]),
        int(lines[1].replace(" ", "").split(":")[1]),
    )
    ways_to_win = get_ways_to_win(race)

    print("Part 2: ", ways_to_win)


def main():
    """Main."""
    part_1()
    part_2()
