"""Main."""
import sys

from twentythree.python.eight import main as eight
from twentythree.python.eleven import main as eleven
from twentythree.python.five import main as five
from twentythree.python.four import main as four
from twentythree.python.fourteen import main as fourteen
from twentythree.python.nine import main as nine
from twentythree.python.seven import main as seven
from twentythree.python.six import main as six
from twentythree.python.ten import main as ten
from twentythree.python.thirteen import main as thirteen
from twentythree.python.twelve import main as twelve

DAYS = {
    4: four.main,
    5: five.main,
    6: six.main,
    7: seven.main,
    8: eight.main,
    9: nine.main,
    10: ten.main,
    11: eleven.main,
    12: twelve.main,
    13: thirteen.main,
    14: fourteen.main,
}

if __name__ == "__main__":
    if len(sys.argv) != 2:  # noqa: PLR2004
        raise Exception("Pass day.")

    DAYS[int(sys.argv[1])]()
