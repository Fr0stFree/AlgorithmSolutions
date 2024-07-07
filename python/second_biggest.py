# Напишите программу, которая принимает на вход N целых чисел и выводит на экран второе по величине число.
# При решении задачи учитывайте, что среди чисел, подающихся на вход программе, могут быть повторяющиеся.


def solve(numbers: list[int]) -> int:
    biggest, second = -1, -1

    for number in filter(lambda num: num > second, numbers):
        if number > biggest:
            biggest, second = number, biggest
        elif biggest > number > second:
            second = number

    return second


def read_data() -> list[int]:
    _ = input()
    return list(map(int, input().strip().split(" ")))


assert solve([4, 3, 5, 2, 7, 1]) == 5
assert solve([4, 7, 4, 2, 7, 1]) == 4
assert solve([5, 4, 3, 2, 1]) == 4
assert solve([1, 2, 3, 4, 5]) == 4

if __name__ == '__main__':
    data = read_data()
    result = solve(data)
    print(result)
