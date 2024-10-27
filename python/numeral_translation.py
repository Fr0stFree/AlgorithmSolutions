"""
Напишите рекурсивную процедуру для перевода десятичного числа в P-ичную систему счисления.
В данной задаче запрещено использовать циклы и массивы.
"""
def translate_recursively(number: int, p: int) -> str:
    if number < p:
        return str(number % p)

    return translate_recursively(number // p, p) + str(number % p)

p, data = int(input()), int(input())
result = translate_recursively(data, p)
print(f"{data}(10)={result}({p})")
