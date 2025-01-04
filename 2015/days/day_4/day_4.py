import hashlib
from pathlib import Path
from typing import Callable

import pytest

FILENAME = "days/day_4/day_4.txt"


def part_1(path: str) -> int:
    with open(path) as f:
        input_data = f.read().strip()
        for i in range(1000000):
            hash_result = hashlib.md5((input_data + str(i)).encode()).hexdigest()
            if hash_result.startswith("00000"):
                return i
    return 0


def part_2(path: str) -> int:
    with open(path) as f:
        input_data = f.read().strip()
        # A very brute force solution
        for i in range(1000000000000000):
            hash_result = hashlib.md5((input_data + str(i)).encode()).hexdigest()
            if hash_result.startswith("000000"):
                return i
    return 0


@pytest.fixture
def temp_file_generator(tmp_path: Path) -> Callable[[str], str]:
    def create_temp_file(content: str) -> str:
        temp_path = tmp_path / "test_input.txt"
        temp_path.write_text(content)
        return str(temp_path)

    return create_temp_file


def test_part_1(temp_file_generator):
    input_content = "abcdef"
    path = temp_file_generator(input_content)
    assert part_1(path) == 609043


if __name__ == "__main__":
    print("Part 1:", part_1(FILENAME))
    print("Part 2:", part_2(FILENAME))
