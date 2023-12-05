"""Parse."""
from pathlib import Path


def read_file(path: Path) -> list[str]:
    """Read file."""
    return [line for line in path.read_text(encoding="utf-8").splitlines()]
