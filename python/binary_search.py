def binary_search(array: list[int], element: int) -> int | None:
    left_idx = 0
    right_idx = len(array) - 1

    while right_idx - left_idx != 1:
        middle_idx = (right_idx + left_idx) // 2
        if array[middle_idx] > element:
            right_idx = middle_idx
        elif array[middle_idx] < element:
            left_idx = middle_idx
        else:
            return middle_idx

    return None
