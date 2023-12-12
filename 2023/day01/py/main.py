"""Advent of code 2023

TODO
----
- [x] Day 01 : part 1
- [ ] Day 01 : part 2

"""
import sys
from pathlib import Path

text_int = dict(
    one = "1",
    two = "2",
    three = "3",
    four = "4",
    five = "5",
    six = "6",
    seven = "7",
    eight = "8",
    nine = "9",
)


def numberfy(text):
    try:
        number = text_int[text]
    except KeyError:
        number = text
    return number



def test_numberfy():
    assert numberfy("one") == "1"
    assert numberfy("two") == "2"
    assert numberfy("three") == "3"
    assert numberfy("four") == "4"
    assert numberfy("five") == "5"
    assert numberfy("six") == "6"
    assert numberfy("seven") == "7"
    assert numberfy("eight") == "8"
    assert numberfy("nine") == "9"
    assert numberfy("10") == "10"
    print("test_numberfy() passed")


def find_number(line, text, digit, out):
    if line == "":
        return out

    if (idx := line.find(text)) >= 0:
        out.append((idx, digit))
        if len(text) > 1:
            new_line = line[:].replace(text[:-1], "-" * len(text[:-1]), 1)
        else:
            new_line = line[:].replace(text, "-", 1)
        find_number(new_line, text, digit, out)
    return out



def test_find_number():
    #      "01234567890123456"
    line = "eight1asdfeightwo"
    out = []
    find_number(line, "eight", "8", out)
    find_number(line, "two", "2", out)
    assert out == [(0, "8"), (10, "8"), (14, "2")]


def solve(line):
    # line = "eight1asdfeightwo"
    print(f"{line = }")
    out = []
    for text, digit in text_int.items():
        out = find_number(line, text, digit, out)
        out = find_number(line, digit, digit, out)

    out = sorted(out, key=lambda x: x[0])
    print(f"{out = }")
    number = int(out[0][1] + out[-1][1])
    print(f"{number = }")
    return number


def main():
    fname = sys.argv[1]
    total = 0
    for i, line in enumerate(Path(fname).read_text().strip().splitlines()):
        number = solve(line)
        total += number
    print(f"{total = }")


if __name__ == "__main__":
    # test_numberfy()
    main()
