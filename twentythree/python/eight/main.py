"""Eight."""
import math
from itertools import cycle
from pathlib import Path
from typing import NamedTuple

from ..utils.parse import read_file

# ruff: noqa: PLR2004

path = Path(__file__).parent.parent.parent / "input" / "8.txt"


class Node(NamedTuple):
    """Node."""

    left: str
    right: str


Network = dict[str, Node]
Directions = list[str]


def parse_input() -> tuple[Network, Directions]:
    """Parse input."""
    lines = read_file(path)
    directions = lines[0]
    network = {}
    for line in lines[2:]:
        name, left_right = line.strip(")").split(" = (")
        network[name] = Node(*left_right.split(", "))

    return network, directions


def part_1():
    """Part 1."""
    network, directions = parse_input()

    i = 0
    cnt = 0
    node = "AAA"
    while node != "ZZZ":
        if i == len(directions):
            i = 0

        if directions[i] == "R":
            node = network[node].right
        else:
            node = network[node].left

        cnt += 1
        i += 1

    print("Part 1: ", cnt)


def part_2():
    """Part 2."""
    network, directions = parse_input()
    nodes = [node for node in network if node[2] == "A"]

    cycles = []
    for node in nodes:
        current = node

        for steps, d in enumerate(cycle(directions), start=1):
            if d == "R":
                current = network[current].right
            else:
                current = network[current].left

            if current[2] == "Z":
                cycles.append(steps)
                break

    print("Part 2: ", math.lcm(*cycles))


def main():
    """Main."""
    part_1()
    part_2()
