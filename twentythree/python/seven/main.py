"""Seven."""
from collections import Counter
from pathlib import Path

from ..utils.parse import read_file

# ruff: noqa: PLR2004

path = Path(__file__).parent.parent.parent / "input" / "7.txt"

STRENGTH = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]
TYPE_STRENGTH = [
    "Five of a kind",
    "Four of a kind",
    "Full house",
    "Three of a kind",
    "Two pair",
    "High card",
]

Hand = str
Rank = int


class Game:
    """A game."""

    hand: str
    bid: int
    type_: int

    def __init__(self, hand: str, bid: str):
        """Init."""
        self.hand = hand
        self.bid = int(bid)
        self.type_ = self._determine_type(hand)

    def _determine_type(self, hand: str) -> int:  # noqa: PLR0911
        """Determine hand."""
        count = Counter(hand)

        if 5 in count.values():
            return 7
        if 4 in count.values():
            return 6
        if 3 in count.values() and len(count) == 2:
            return 5
        if 3 in count.values():
            return 4
        if 2 in count.values() and len(count) == 3:
            return 3
        if 2 in count.values() and len(count) == 4:
            return 2
        if len(count) == 5:
            return 1

        raise Exception("Should not be here.")


def parse_input() -> list[Game]:
    """Parse input."""
    return [Game(line.split()[0], int(line.split()[1])) for line in read_file(path)]


def _is_stronger(left: Game, right: Game) -> bool:
    """Is "left" stronger than "right"."""
    if left.type_ > right.type_:
        return True

    if left.type_ < right.type_:
        return False

    for i in range(len(left.hand)):
        left_card = STRENGTH.index(left.hand[i])
        right_card = STRENGTH.index(right.hand[i])

        if left_card > right_card:
            return True
        if left_card < right_card:
            return False

    return False


def _insertion_sort(games: list[Game]) -> list[Game]:
    """Do a insertion sort of the games."""
    i = 1
    while i < len(games):
        j = i
        while j > 0:
            if _is_stronger(games[j - 1], games[j]):
                games.insert(j - 1, games.pop(j))
            j -= 1

        i += 1

    return games


def part_1():
    """Part 1."""
    games = parse_input()
    games = _insertion_sort(games)
    total = 0
    for rank, g in enumerate(games):
        total += (rank + 1) * g.bid
    print("Part 1: ", total)


def main():
    """Main."""
    part_1()
