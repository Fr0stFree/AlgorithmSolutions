def compute_sum(array: list[int]) -> int:
    if len(array) == 0:
        return 0
    return array[0] + compute_sum(array[1:])

_, array = input(), list(map(int, input().split(" ")))
result = compute_sum(array)
print(result)