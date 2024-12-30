FILENAME = "days/day_2/day_2.txt"


def part_1(path: str) -> int:
    with open(path) as f:
        content = f.read()
    print(content)
    return 0


def part_2(path: str) -> int:
    with open(path) as f:
        content = f.read()
    print(content)
    floor = 0
    return floor


if __name__ == "__main__":
    print(part_1(FILENAME))
    print(part_2(FILENAME))
