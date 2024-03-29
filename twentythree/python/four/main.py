"""Four."""
from collections import Counter
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "4.txt"

Numbers = list[int]

Card = tuple[Numbers, Numbers]


def parse_input(lines: list[str]) -> list[Card]:
    """Parse the cards."""
    cards: list[Card] = []
    for line in lines:
        split: list[str] = line.split(":")[1].split("|")
        winning = [int(nbr) for nbr in split[0].strip().split(" ") if nbr != ""]
        nbrs = [int(nbr) for nbr in split[1].strip().split(" ") if nbr != ""]
        cards.append((winning, nbrs))

    return cards


def part_1(cards: list[Card]):
    """Calculate worth of cards."""
    worth = 0
    for card in cards:
        winning = [nbr for nbr in card[1] if nbr in card[0]]
        exponent = len(winning) - 1
        if exponent >= 0:
            worth += pow(2, exponent)

    print("Part 1: ", worth)


def part_2(cards: list[Card]):
    """Calculate number of copies."""
    count = Counter()
    # Iterate cards
    for i, card in enumerate(cards):
        matching = len([nbr for nbr in card[1] if nbr in card[0]])

        # Add one copy
        count.update([i])

        # For each copy
        for _ in range(count.get(i)):
            # Add copies of "matching" number of next cards
            for j in range(i + 1, i + matching + 1):
                count.update([j])

    print("Part 2: ", count.total())


def main():
    """Main."""
    cards = parse_input(read_file(path))
    part_1(cards)
    part_2(cards)
