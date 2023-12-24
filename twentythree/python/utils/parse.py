"""Parse."""
from pathlib import Path


def read_file(path: Path) -> list[str]:
    """Read file."""
    lines = [line for line in path.read_text(encoding="utf-8").splitlines()]
    if lines[-1] == "":
        lines = lines[:-1]
    return lines
