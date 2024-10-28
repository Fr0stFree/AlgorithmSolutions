def find_max_element(array: list[int]) -> int:
    if len(array) == 1:
        return array[0]

    next_max = find_max_element(array[1:])
    return max(next_max, array[0])


_, array = input(), list(map(int, input().split(" ")))
result = find_max_element(array)
print(result)