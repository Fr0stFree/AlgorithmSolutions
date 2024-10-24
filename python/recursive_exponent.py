"""
Напишите рекурсивную функцию calculatePower, которая вычисляет степень числа. Функция calculatePower должна принимать
два аргумента: число base и целое число exponent. Она должна возвращать результат возведения числа base в степень exponent.
"""


def calculatePower(base: int, exp: int, result = 1) -> float:
    if exp == 0:
        return result
    result *= base
    exp -= 1
    return calculatePower(base, exp, result)

data = map(int, input().split(" "))
result = calculatePower(*data)
print(f"{result:.2f}")
