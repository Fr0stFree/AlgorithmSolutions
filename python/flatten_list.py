def flatten(data: list[int | list]) -> list[int]:
    result = []

    for element in data:
        if isinstance(element, int):
            result.append(element)
        else:
            data.extend(element)

    return result

assert flatten([3, 1, [7, [9, 1]], 4, [7], [[9, [7]], 1, 2], 3, [3, 5, 1, 6]]) == [3, 1, 7, 9, 1, 4, 7, 9, 7, 1, 2, 3, 3, 5, 1, 6]