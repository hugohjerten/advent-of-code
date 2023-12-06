"""Five."""
from pathlib import Path
from typing import NamedTuple

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "5.txt"


class Map(NamedTuple):
    """A map."""

    destination: int
    source: int
    range_: int

    def get_destination(self, nbr: int) -> int | None:
        """Map a nbr; return None if not in range."""
        if nbr >= self.source and nbr <= self.source + self.range_ - 1:
            return nbr - self.source + self.destination
        return None


MapType = list[Map]


class Maps:
    """All maps."""

    types: list[MapType]

    def __init__(self, types: list[MapType]):
        """All maps."""
        self.types = types

    def __map_type(self, source: int, type_: MapType) -> int:
        """Map destination for source."""
        for map_ in type_:
            dest = map_.get_destination(source)
            if dest:
                return dest

        return source

    def get_location(self, seed: int) -> int:
        """Get location for seed."""
        source = seed
        for type_ in self.types:
            source = self.__map_type(source, type_)

        return source


def parse_input(lines: list[str]) -> tuple[list[int], Maps]:
    """Parse input."""
    seeds = [int(seed) for seed in lines[0].split(" ") if seed.isdigit()]
    maps: list[MapType] = []
    current = 0

    for line in lines[2:]:
        if line == "":
            current += 1
            continue

        if "map" in line:
            maps.append([])
            continue

        nbrs = line.split(" ")
        maps[current].append(
            Map(destination=int(nbrs[0]), source=int(nbrs[1]), range_=int(nbrs[2]))
        )

    return seeds, Maps(maps)


def part1(seeds: list[int], maps: Maps):
    """Part 1."""
    locations = []
    for seed in seeds:
        locations.append(maps.get_location(seed))

    print("Part 1: ", min(locations))


def main():
    """Main."""
    seeds, maps = parse_input(read_file(path))
    part1(seeds, maps)
