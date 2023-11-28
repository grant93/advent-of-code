#!/bin/python3

import sys
import functools

sys.setrecursionlimit(100000)

FILE = sys.argv[1] if len(sys.argv) > 1 else "input.txt"


def decrypt(arrangement, key=1, num_mix=1):
    zero_index = 0
    mixer = [(itx, a * key) for (itx, a) in enumerate(arrangement)]

    for _ in range(num_mix):
        for i in range(len(arrangement)):
            for j in range(len(arrangement)):
                if mixer[j][0] == i:
                    curr = mixer.pop(j)
                    new_index = (j + curr[1]) % len(mixer)
                    mixer.insert(new_index, (i, curr[1]))
                    break

    for (itx, (_, val)) in enumerate(mixer):
        if val == 0:
            zero_index = itx
            break

    print(mixer)
    return (
        mixer[(zero_index + 1000) % len(mixer)][1]
        + mixer[(zero_index + 2000) % len(mixer)][1]
        + mixer[(zero_index + 3000) % len(mixer)][1]
    )


def main():
    print(f"Using file {FILE}")
    with open(FILE, "r", encoding="utf-8") as f:
        arrangement = [int(x) for x in f.readlines()]

        print(f"Part one: {decrypt(arrangement.copy())}")
        print(f"Part one: {decrypt(arrangement.copy(), 811589153, 10)}")


main()

