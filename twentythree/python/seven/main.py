"""Seven."""
from collections import Counter
from pathlib import Path

from ..utils.parse import read_file

# ruff: noqa: PLR2004

path = Path(__file__).parent.parent.parent / "input" / "7.txt"

STRENGTH = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]
STRENGTH_PART_2 = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"]
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

    def __init__(self, hand: str, bid: str, part_2: bool = False):
        """Init."""
        self.hand = hand
        self.bid = int(bid)
        self.type_ = (
            self._determine_type_2(hand) if part_2 else self._determine_type(hand)
        )

    def _determine_type_2(self, hand: str) -> int:  # noqa: PLR0911
        """Determine type part 2."""
        if "J" not in hand:
            return self._determine_type(hand)

        jokers = Counter(hand)["J"]
        count = Counter(hand.replace("J", ""))

        if jokers in (4, 5) or (jokers in (1, 2, 3) and len(count) == 1):
            return 7
        if (jokers in (2, 3) and len(count) == 2) or (
            len(count) == 2 and 3 in count.values() and jokers == 1
        ):
            return 6
        if len(count) == 2 and jokers == 1:
            return 5
        if (len(count) == 3 and jokers == 1) or (len(count) == 3 and jokers == 2):
            return 4
        if len(count) == 3 and jokers == 1:
            return 3
        if len(count) == 4 and jokers == 1:
            return 2

        raise Exception("Should not be here.")

    def _determine_type(self, hand: str) -> int:  # noqa: PLR0911
        """Determine type."""
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


def parse_input(part_2: bool = False) -> list[Game]:
    """Parse input."""
    return [
        Game(line.split()[0], int(line.split()[1]), part_2) for line in read_file(path)
    ]


def _is_stronger(left: Game, right: Game, part_2: bool) -> bool:
    """Is "left" stronger than "right"."""
    if left.type_ > right.type_:
        return True

    if left.type_ < right.type_:
        return False

    strength = STRENGTH_PART_2 if part_2 else STRENGTH

    for i in range(len(left.hand)):
        left_card = strength.index(left.hand[i])
        right_card = strength.index(right.hand[i])

        if left_card > right_card:
            return True
        if left_card < right_card:
            return False

    return False


def _insertion_sort(games: list[Game], part_2: bool = False) -> list[Game]:
    """Do a insertion sort of the games."""
    i = 1
    while i < len(games):
        j = i
        while j > 0:
            if _is_stronger(games[j - 1], games[j], part_2):
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


def part_2():
    """Part 2."""
    part_2 = True
    games = parse_input(part_2)
    games = _insertion_sort(games, part_2)
    total = 0
    for rank, g in enumerate(games):
        total += (rank + 1) * g.bid
    print("Part 2: ", total)


def main():
    """Main."""
    part_1()
    part_2()
