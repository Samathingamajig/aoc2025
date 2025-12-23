import os.path
import z3

BASE_PATH = os.path.join("day10")
EXAMPLE_INPUT_PATH = os.path.join(BASE_PATH, "example.input.txt")
EXAMPLE_OUTPUT_PATH = os.path.join(BASE_PATH, "b.example.output.txt")
REAL_INPUT_PATH = os.path.join(BASE_PATH, "input.txt")


def solution(inp: str, is_example: bool) -> str:
    result = 0
    for line in inp.splitlines():
        groups = line.split(" ")

        buttons = []
        for buttonRaw in groups[1:-1]:
            buttons.append([int(n) for n in buttonRaw[1:-1].split(",")])

        joltageReq = [int(n) for n in groups[-1][1:-1].split(",")]

        opt = z3.Optimize()

        knobs = [z3.Int(f"k{i}") for i in range(len(buttons))]

        for knob in knobs:
            opt.add(knob >= 0)

        for i, targetJoltage in enumerate(joltageReq):
            sumEq = 0
            for buttonIdx, button in enumerate(buttons):
                if any(idx == i for idx in button):
                    sumEq += knobs[buttonIdx]
            opt.add(sumEq == targetJoltage)

        opt.minimize(sum(knobs))

        if opt.check() != z3.sat:
            return "error"

        model = opt.model()
        result += sum(model[knob].as_long() for knob in knobs)

    return str(result)


def main():
    with open(EXAMPLE_INPUT_PATH, "r") as f:
        exampleInput = f.read().rstrip("\n")
    with open(EXAMPLE_OUTPUT_PATH, "r") as f:
        exampleOutput = f.read().rstrip("\n")
    with open(REAL_INPUT_PATH, "r") as f:
        realInput = f.read().rstrip("\n")

    print("Running example...")
    exampleAttempt = solution(exampleInput, True)

    if exampleAttempt == exampleOutput:
        print(f"✅ Example correct! `{exampleAttempt}`")
    else:
        print("❌ Example incorrect")
        print(f"Expected: `{exampleOutput}`")
        print(f"Received: `{exampleAttempt}`")
        return 1

    print("Running real...")

    realAttempt = solution(realInput, False)

    print(f"Calculated result: `{realAttempt}`")


if __name__ == "__main__":
    raise SystemExit(main())
