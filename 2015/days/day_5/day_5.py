from pathlib import Path
from typing import Callable

import pytest

FILENAME = "days/day_5/day_5.txt"


def part_1(path: str) -> int:
    total = 0
    with open(path) as f:
        lines = f.read().strip().split("\n")
        for line in lines:
            vowel_count = 0
            has_double_letter = False
            has_forbidden_string = False
            for i, c in enumerate(line):
                if c in "aeiou":
                    vowel_count += 1
                # skip check first element for double letter or forbidden string
                if i == 0:
                    continue

                c_prev = line[i - 1]
                if c == c_prev:
                    has_double_letter = True

                if c_prev + c in ("ab", "cd", "pq", "xy"):
                    has_forbidden_string = True
                    break

            if vowel_count >= 3 and has_double_letter and not has_forbidden_string:
                total += 1
    return total


def part_2(path: str) -> int:
    with open(path) as f:
        input_data = f.read().strip()
    return 0


@pytest.fixture
def temp_file_generator(tmp_path: Path) -> Callable[[str], str]:
    def create_temp_file(content: str) -> str:
        temp_path = tmp_path / "test_input.txt"
        temp_path.write_text(content)
        return str(temp_path)

    return create_temp_file


def test_part_1(temp_file_generator):
    input_content = (
        "ugknbfddgicrmopn\naaa\njchzalrnumimnmhp\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb"
    )
    path = temp_file_generator(input_content)
    assert part_1(path) == 2


if __name__ == "__main__":
    print("Part 1:", part_1(FILENAME))
    print("Part 2:", part_2(FILENAME))
