"""
Напишите рекурсивную функцию calculatePower, которая вычисляет степень числа. Функция calculatePower должна принимать
два аргумента: число base и целое число exponent. Она должна возвращать результат возведения числа base в степень exponent.
"""
def calculate_power(base: int, exp: int) -> float:
    if exp == 1:
        return base
    return base * calculate_power(base, exp-1)

if __name__ == '__main__':
    data = map(int, input().split(" "))
    result = calculate_power(*data)
    print(result)
