from functools import lru_cache

@lru_cache
def calculator(number: int) -> list[int]:
    if number < 1:
        return []

    options = []

    if number % 2 == 0:
        option = [number] + calculator(number // 2)
        options.append(option)

    if number % 3 == 0:
        option = [number] + calculator(number // 3)
        options.append(option)

    option = [number] + calculator(number - 1)
    options.append(option)

    best_option = min(options, key=len)
    return best_option



if __name__ == "__main__":
    initial_number = 83
    result = calculator(initial_number)
    print(f" -> ".join(map(str, result)))
