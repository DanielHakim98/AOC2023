import pytest

FILENAME = "days/day_1/day_1.txt"


def part_1(path: str) -> int:
    with open(path) as f:
        content = f.read()

    floor = 0
    for c in content:
        if c == "(":
            floor += 1
        else:
            floor -= 1
    return floor


def part_2(path: str) -> int:
    with open(path) as f:
        content = f.read()

    floor = 0
    for i, c in enumerate(content):
        # print(i, c)
        if c == "(":
            floor += 1
        else:
            floor -= 1

        if floor == -1:
            return i + 1

    return floor


@pytest.fixture
def temp_file(tmp_path):
    def create_temp_file(content):
        temp_path = tmp_path / "test_input.txt"
        temp_path.write_text(content)
        return str(temp_path)

    return create_temp_file


def test_part_1(temp_file):
    input_content = "(())"
    path = temp_file(input_content)
    assert part_1(path) == 0

    input_content = "(()(()("
    path = temp_file(input_content)
    assert part_1(path) == 3

    input_content = "))("
    path = temp_file(input_content)
    assert part_1(path) == -1

    input_content = ")())())"
    path = temp_file(input_content)
    assert part_1(path) == -3


def test_part_2(temp_file):
    input_content = "((())))"
    path = temp_file(input_content)
    assert part_2(path) == 7


if __name__ == "__main__":
    print(part_1(FILENAME))
    print(part_2(FILENAME))
