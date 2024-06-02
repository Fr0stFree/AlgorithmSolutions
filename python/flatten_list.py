def flatten(data: list[int | list]) -> list[int]:
    result = []
    todo = []

    for element in data:
        todo.append(element)

        while todo:
            current = todo.pop()

            if isinstance(current, int):
                result.append(current)
            else:
                for item in reversed(current):
                    todo.append(item)

    return result


assert flatten([3, 1, [7, [9, 1]], 4, [7], [[9, [7]], 1, 2], 3, [3, 5, 1, 6]]) == [3, 1, 7, 9, 1, 4, 7, 9, 7, 1, 2, 3, 3, 5, 1, 6]