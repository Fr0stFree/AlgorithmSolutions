def merge_sort(array: list[int]) -> list[int]:
    if len(array) == 1:
        return array

    left_half = array[:len(array) // 2]
    right_half = array[len(array) // 2:]
    sorted_left_half = merge_sort(left_half)
    sorted_right_half = merge_sort(right_half)
    return merge(sorted_left_half, sorted_right_half)


def merge(left_array: list[int], right_array: list[int]) -> list[int]:
    left_idx = 0
    right_idx = 0
    result = []

    while True:
        if left_array[left_idx] <= right_array[right_idx]:
            result.append(left_array[left_idx])
            left_idx += 1
            if left_idx == len(left_array):
                result.extend(right_array[right_idx:])
                break
        else:
            result.append(right_array[right_idx])
            right_idx += 1
            if right_idx == len(right_array):
                result.extend(left_array[left_idx:])
                break

    return result
