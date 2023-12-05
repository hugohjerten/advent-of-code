"""Main."""
import sys

from twentythree.python.four import main as four

DAYS = {4: four.main}

if __name__ == "__main__":
    if len(sys.argv) != 2:  # noqa: PLR2004
        raise Exception("Pass day.")

    DAYS[int(sys.argv[1])]()
