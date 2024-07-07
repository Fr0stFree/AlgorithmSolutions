# У вас есть массив целых чисел a. Вам необходимо ответить на q запросов.
# В i-м запросе вам дают число xᵢ и спрашивают: «Какое число в массиве имеет значение, максимально близкое к xᵢ?»
# Если подходящих чисел несколько, в качестве ответа на запрос выведите наименьшее.
import math


def solve(digits: list[int], looking_digits: list[int]) -> list[int]:
    result: list[int] = []
    digits.sort()
    offset = 0

    for looking_digit in looking_digits:
        value, index = search_nearest(looking_digit, digits[offset:])
        offset = index
        result.append(value)

    return result


def search_nearest(looking_digit: int, digits: list[int]) -> (int, int):
    prev_diff = math.inf
    nearest_digit = digits[0]
    current_index = 0

    for index, digit in enumerate(digits):
        new_diff = abs(looking_digit - digit)
        if new_diff > prev_diff:
            break

        nearest_digit = digit if new_diff != prev_diff else min(digit, nearest_digit)
        prev_diff = new_diff
        current_index = index

    return nearest_digit, current_index


def read_data() -> (list[int], list[int]):
    to_int_array = lambda raw: list(map(int, raw.strip().split(" ")))
    _ = input()
    digits = to_int_array(input())
    _ = input()
    looking_digits = to_int_array(input())
    return digits, looking_digits


assert solve([9, 4, 8, 1, 4], [1, 3, 6, 9, 10]) == [1, 4, 4, 9, 9]


if __name__ == '__main__':
    data = read_data()
    result = solve(*data)
    print(result)
