from pathlib import Path
from typing import Callable

import pytest

FILENAME = "days/day_3/day_3.txt"


def part_1(path: str) -> int:
    house: set[tuple[int, int]] = set()
    house.add((0, 0))
    with open(path) as f:
        content = f.read()
        x, y = 0, 0
        for c in content:
            if c == ">":
                x += 1
            elif c == "<":
                x -= 1
            elif c == "^":
                y += 1
            elif c == "v":
                y -= 1
            house.add((x, y))
    return len(house)


def part_2(path: str) -> int:
    houses: set[tuple[int, int]] = set()
    houses.add((0, 0))
    with open(path) as f:
        content = f.read()
        x1, y1 = 0, 0
        x2, y2 = 0, 0
        is_robot = False
        for c in content:
            if c == ">":
                if is_robot:
                    x2 += 1
                else:
                    x1 += 1
            elif c == "<":
                if is_robot:
                    x2 -= 1
                else:
                    x1 -= 1
            elif c == "^":
                if is_robot:
                    y2 += 1
                else:
                    y1 += 1
            elif c == "v":
                if is_robot:
                    y2 -= 1
                else:
                    y1 -= 1
            house = (x1, y1) if not is_robot else (x2, y2)
            houses.add(house)
            is_robot = not is_robot

    return len(houses)


@pytest.fixture
def temp_file_generator(tmp_path: Path) -> Callable[[str], str]:
    def create_temp_file(content: str) -> str:
        temp_path = tmp_path / "test_input.txt"
        temp_path.write_text(content)
        return str(temp_path)

    return create_temp_file


def test_part_1_test_1(temp_file_generator):
    input_content = ">"
    path = temp_file_generator(input_content)
    assert part_1(path) == 2


def test_part_1_test_2(temp_file_generator):
    input_content = "^>v<"
    path = temp_file_generator(input_content)
    assert part_1(path) == 4


def test_part_1_test_3(temp_file_generator):
    input_content = "^v^v^v^v^v"
    path = temp_file_generator(input_content)
    assert part_1(path) == 2


def test_part_2_test_1(temp_file_generator):
    input_content = "^v"
    path = temp_file_generator(input_content)
    assert part_2(path) == 3


def test_part_2_test_2(temp_file_generator):
    input_content = "^>v<"
    path = temp_file_generator(input_content)
    assert part_2(path) == 3


def test_part_2_test_3(temp_file_generator):
    input_content = "^v^v^v^v^v"
    path = temp_file_generator(input_content)
    assert part_2(path) == 11


if __name__ == "__main__":
    print(part_1(FILENAME))
    print(part_2(FILENAME))
