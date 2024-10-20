"""
Дано целое число 𝑛 (1≤𝑛≤1012). Нужно найти натуральные 𝑎,𝑏,𝑐:𝑎𝑏𝑐=𝑛 и при этом 2(𝑎𝑏+𝑏𝑐+𝑐𝑎) минимально.
Т.е. при фиксированном объеме минимимизировать площадь поверхности.
"""
import math


def solve(volume: int) -> tuple[int, int, int, int]:
    smallest_area = math.inf
    sides: list[int] = [0] * 3

    a = 0
    while a*a*a < volume:
        a += 1
        b = 0
        while b*b*a < volume:
            b += 1
            c = volume / a / b
            if not c.is_integer():
                continue

            c = int(c)
            area = int(2 * (a*b + a*c + b*c))

            if area < smallest_area:
                sides[0], sides[1], sides[2] = a, b, c
                smallest_area = area

    return smallest_area, sides[0], sides[1], sides[2]





if __name__ == '__main__':
    volume = int(input())
    result = solve(volume)
    print(" ".join(map(str, result)))