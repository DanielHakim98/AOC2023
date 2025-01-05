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
    total = 0
    with open(path) as f:
        lines = f.read().strip().split("\n")
        # At first, I thought that I need to implement sliding window to solve this
        # but turns out I can just do it this way
        for line in lines:
            has_double_pair = False
            has_triplet = False
            # to store the index of the first pair, later can be used to check the interval
            pairs: dict[str, int] = {}
            for i in range(1, len(line)):
                # Check for double pair
                pair = line[i - 1] + line[i]
                if pair in pairs:
                    start = pairs[pair]
                    interval = [start, start + 1]
                    if interval[1] != i - 1:
                        has_double_pair = True
                else:
                    pairs[pair] = i - 1

                # Check for triplet with 1 distinct letter at middle
                if i > 1:
                    triplet = line[i - 2] + line[i - 1] + line[i]
                    if triplet[0] == triplet[2] and triplet[0] != triplet[1]:
                        has_triplet = True

            if has_double_pair and has_triplet:
                total += 1
    return total


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


def test_part_2(temp_file_generator):
    input_content = "qjhvhtzxzqqjkmpb\nxxyxx\nuurcxstgmygtbstg\nieodomkazucvgmuy"
    path = temp_file_generator(input_content)
    assert part_2(path) == 2


if __name__ == "__main__":
    print("Part 1:", part_1(FILENAME))
    print("Part 2:", part_2(FILENAME))
