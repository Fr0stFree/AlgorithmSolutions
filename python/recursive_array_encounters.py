def find_encounters(array: list[int], value: int) -> int:
    if len(array) == 0:
        return 0

    counter = 1 if value == array[0] else 0
    return counter + find_encounters(array[1:], value)


_, array, looking_value = input(), list(map(int, input().split(" "))), int(input())
result = find_encounters(array, looking_value)
print(result)