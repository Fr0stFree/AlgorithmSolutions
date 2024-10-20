"""
Дано целое число 𝑛 (1≤𝑛≤1012). Нужно найти натуральные 𝑎,𝑏,𝑐:𝑎𝑏𝑐=𝑛 и при этом 2(𝑎𝑏+𝑏𝑐+𝑐𝑎) минимально.
Т.е. при фиксированном объеме минимимизировать площадь поверхности.
"""
import math


def solve(volume: int) -> tuple[int, int, int, int]:
    smallest_area = math.inf
    sides = [0, 0, 0]

    for a in range(1, int(volume ** (1 / 3)) + 1):
        if volume % a != 0:
            continue

        for b in range(a, int(math.sqrt(volume // a)) + 1):
            if (volume // a) % b != 0:
                continue

            c = volume // (a * b)
            area = 2 * (a * b + b * c + a * c)

            if area < smallest_area:
                smallest_area = area
                sides = sorted([a, b, c])

    return smallest_area, *sides


if __name__ == '__main__':
    volume = int(input())
    result = solve(volume)
    print(" ".join(map(str, result)))