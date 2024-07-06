# Напишите программу, которая принимает на вход N целых чисел и выводит на экран второе по величине число.
# При решении задачи учитывайте, что среди чисел, подающихся на вход программе, могут быть повторяющиеся.


def solve(data: list[int]) -> int:
    biggest, second = -1, -1

    def update_biggest(number: int) -> None:
        nonlocal biggest, second
        second = biggest
        biggest = number

    [update_biggest(element) for element in data if element > biggest]

    return second


def read_data() -> list[int]:
    _ = input()
    return list(map(int, input().strip().split(" ")))


if __name__ == '__main__':
    data = read_data()
    result = solve(data)
    print(result)
