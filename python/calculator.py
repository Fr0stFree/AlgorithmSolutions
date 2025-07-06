def calculator(init_number: int) -> list[int]:
    memo = {}

    def _calculate(number: int) -> list[int]:
        if number < 1:
            return []

        if number in memo:
            return memo[number]

        options = []

        if number % 2 == 0:
            option = [number] + _calculate(number // 2)
            options.append(option)

        if number % 3 == 0:
            option = [number] + _calculate(number // 3)
            options.append(option)

        option = [number] + _calculate(number - 1)
        options.append(option)

        best_option = min(options, key=len)
        memo[number] = best_option
        return memo[number]

    return _calculate(init_number)


if __name__ == "__main__":
    initial_number = 83
    result = calculator(initial_number)
    print(f" -> ".join(map(str, result)))
