"""Nine."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "9.txt"


History = list[int]


def parse_input() -> list[History]:
    """Parse input."""
    return [[int(val) for val in line.split()] for line in read_file(path)]


def _find_next_sequences(history: History) -> list[History]:
    """Find next sequences."""
    seqs = [history]
    while True:
        seqs.append([])
        for i in range(1, len(seqs[-2])):
            seqs[-1].append(seqs[-2][i] - seqs[-2][i - 1])

        if all(val == 0 for val in seqs[-1]):
            break

    return seqs


def _extrapolate_right(history: History) -> int:
    """Extrapolate next value."""
    # Find next sequences
    seqs = _find_next_sequences(history)

    # Append zero at end of last sequence
    seqs[-1].append(0)

    # Extrapolate value for each sequence
    for i in range(len(seqs) - 1)[::-1]:
        seqs[i].append(seqs[i + 1][-1] + seqs[i][-1])

    return seqs[0][-1]


def _extrapolate_left(history: History) -> int:
    """Extrapolate next value."""
    # Find next sequences
    seqs = _find_next_sequences(history)

    # Add zero at beginning of last sequence
    seqs[-1].insert(0, 0)

    # Extrapolate value for each sequence
    for i in range(len(seqs) - 1)[::-1]:
        seqs[i].insert(0, seqs[i][0] - seqs[i + 1][0])

    return seqs[0][0]


def part_1():
    """Part 1."""
    histories = parse_input()
    extrapolated = [_extrapolate_right(h) for h in histories]
    print("Part 1: ", sum(extrapolated))


def part_2():
    """Part 2."""
    histories = parse_input()
    extrapolated = [_extrapolate_left(h) for h in histories]
    print("Part 2: ", sum(extrapolated))


def main():
    """Main."""
    part_1()
    part_2()
