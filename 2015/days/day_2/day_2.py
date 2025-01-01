FILENAME = "days/day_2/day_2.txt"


def part_1(path: str) -> int:
    total = 0
    with open(path) as f:
        content = f.readlines()

        for line in content:
            dimensions = [int(item) for item in line.strip().split("x")]
            leng = dimensions[0]
            widt = dimensions[1]
            heig = dimensions[2]

            leng_widt = leng * widt
            widt_heig = widt * heig
            heig_leng = heig * leng
            min_surface = min(leng_widt, widt_heig, heig_leng)
            area = 2 * leng_widt + 2 * widt_heig + 2 * heig_leng + min_surface
            total += area
    return total


def part_2(path: str) -> int:
    total = 0
    with open(path) as f:
        content = f.readlines()
        for line in content:
            dimensions = [int(item) for item in line.strip().split("x")]
            leng = dimensions[0]
            widt = dimensions[1]
            heig = dimensions[2]

            leng_widt = 2 * leng + 2 * widt
            widt_heig = 2 * widt + 2 * heig
            heig_leng = 2 * heig + 2 * leng
            min_peri = min(leng_widt, widt_heig, heig_leng)
            ribbon_feet = min_peri + leng * widt * heig
            total += ribbon_feet

    return total


if __name__ == "__main__":
    print(part_1(FILENAME))
    print(part_2(FILENAME))
