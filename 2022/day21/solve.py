#!/bin/python3

import sys

sys.setrecursionlimit(100000)

FILE = sys.argv[1] if len(sys.argv) > 1 else "input.txt"


def part_one(monkeys):
    done = dict()

    while "root" not in done:
        for monkey in monkeys:
            (name, job) = monkey
            if name not in done:
                if len(job) == 1:
                    done[name] = job[0]
                elif job[0] in done and job[2] in done:
                    if job[1] == "+":
                        done[name] = done[job[0]] + done[job[2]]
                    elif job[1] == "-":
                        done[name] = done[job[0]] - done[job[2]]
                    elif job[1] == "*":
                        done[name] = done[job[0]] * done[job[2]]
                    elif job[1] == "/":
                        done[name] = done[job[0]] // done[job[2]]

    return done["root"]


def part_two(monkeys):
    monkeys = dict(monkeys)
    monkeys["root"][1] = "="
    del monkeys["humn"]

    def build_ast(curr):
        if curr == "humn":
            return ["x"]

        parts = monkeys[curr]
        if len(parts) == 1:
            return parts
        else:
            left_child = build_ast(parts[0])
            right_child = build_ast(parts[2])

            if (
                len(left_child) == 1
                and len(right_child) == 1
                and isinstance(left_child[0], int)
                and isinstance(right_child[0], int)
            ):
                lhs = left_child[0]
                rhs = right_child[0]

                if parts[1] == "+":
                    return [int(lhs) + int(rhs)]
                elif parts[1] == "-":
                    return [int(lhs) - int(rhs)]
                elif parts[1] == "*":
                    return [int(lhs) * int(rhs)]
                elif parts[1] == "/":
                    return [int(lhs) // int(rhs)]
            else:
                return [left_child, parts[1], right_child]

    ast = build_ast("root")
    lhs = ast[0]
    rhs = ast[2]

    if len(lhs) == 1:
        to_solve = rhs
        other = lhs
    else:
        to_solve = lhs
        other = rhs

    # Now basically just solve for x
    def traverse_tree(curr, solved_side):
        if len(curr) == 1:
            return solved_side
        else:
            [l, op, r] = curr
            if len(l) == 1 and isinstance(l[0], int):
                # move the right side
                match op:
                    case "+":
                        return traverse_tree(r, [solved_side, "-", l])
                    case "-":
                        return traverse_tree(r, [l, "-", solved_side])
                    case "*":
                        return traverse_tree(r, [solved_side, "/", l])
                    case "/":
                        return traverse_tree(r, [l, "*", solved_side])
            elif len(r) == 1 and isinstance(r[0], int):
                # move the left side
                match op:
                    case "+":
                        return traverse_tree(l, [solved_side, "-", r])
                    case "-":
                        return traverse_tree(l, [solved_side, "+", r])
                    case "*":
                        return traverse_tree(l, [solved_side, "/", r])
                    case "/":
                        return traverse_tree(l, [solved_side, "*", r])

    new_side = traverse_tree(to_solve, other)

    def solve_tree(curr):
        if len(curr) == 1:
            return curr
        else:
            left_child = solve_tree(curr[0])
            right_child = solve_tree(curr[2])

            if (
                len(left_child) == 1
                and len(right_child) == 1
                and isinstance(left_child[0], int)
                and isinstance(right_child[0], int)
            ):
                lhs = left_child[0]
                rhs = right_child[0]

                if curr[1] == "+":
                    return [int(lhs) + int(rhs)]
                elif curr[1] == "-":
                    return [int(lhs) - int(rhs)]
                elif curr[1] == "*":
                    return [int(lhs) * int(rhs)]
                elif curr[1] == "/":
                    return [int(lhs) // int(rhs)]
            else:
                return [left_child, curr[1], right_child]

    return solve_tree(new_side)[0]


def main():
    print(f"Using file {FILE}")
    with open(FILE, "r", encoding="utf-8") as f:
        monkeys = []
        for line in f:
            words = line.strip().split(" ")
            name = words[0].split(":")[0]
            job = words[1:]
            if len(job) == 1:
                job = [int(job[0])]
            monkeys.append((name, job))

        print(f"Part one: {part_one(monkeys)}")
        print(f"Part two: {part_two(monkeys)}")


main()

