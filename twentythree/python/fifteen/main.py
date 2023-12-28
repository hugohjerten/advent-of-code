"""Fifteen."""
from pathlib import Path

from ..utils.parse import read_file

path = Path(__file__).parent.parent.parent / "input" / "15.txt"

# Label, focal length
Lens = tuple[str, int]
# key: box number, value: list of lenses
Boxes = dict[int, list[Lens]]


def parse_input() -> list[str]:
    """Parse input."""
    return read_file(path)[0].split(",")


def hash_algorithm(step: str) -> int:
    """Run hash algorithm on step."""
    val = 0
    for c in step:
        val += ord(c)
        val *= 17
        val = val % 256

    return val


def part_1():
    """Part 1."""
    sequence = parse_input()
    total = sum([hash_algorithm(c) for c in sequence])

    print("Part 1: ", total)


def _remove_lens(label: str, box: int, boxes: Boxes) -> Boxes:
    """Remove lens."""
    if box not in boxes:
        return boxes

    for i, lens in enumerate(boxes[box]):
        if label == lens[0]:
            boxes[box].pop(i)

            if not boxes[box]:
                del boxes[box]

            return boxes

    return boxes


def _add_lens(new: Lens, box: int, boxes: Boxes) -> Boxes:
    """Add lens."""
    if box not in boxes:
        boxes[box] = [new]
        return boxes

    for i, lens in enumerate(boxes[box]):
        if new[0] == lens[0]:
            boxes[box][i] = new
            return boxes

    boxes[box].append(new)
    return boxes


def part_2():
    """Part 2."""
    sequence = parse_input()

    boxes: Boxes = {}
    for s in sequence:
        label = s.split("-")[0] if "-" in s else s.split("=")[0]
        box = hash_algorithm(label)

        if "-" in s:
            boxes = _remove_lens(label, box, boxes)
        elif "=" in s:
            boxes = _add_lens(Lens((label, int(s[-1]))), box, boxes)

    total = sum(
        [
            sum([(box + 1) * (i + 1) * lens[1] for i, lens in enumerate(lenses)])
            for box, lenses in boxes.items()
        ]
    )
    print("Part 2:", total)


def main():
    """Main."""
    part_1()
    part_2()
