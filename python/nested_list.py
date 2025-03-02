from typing import Any


def flatten_rec(input: list[Any]) -> list[Any]:
    if not input:
        return []

    elem = input.pop(0)
    if isinstance(elem, list):
        return flatten_rec(elem) + flatten_rec(input)

    return [elem] + flatten_rec(input)


def flatten_queue(queue: list[Any]) -> list[Any]:
    result = []

    while queue:
        item = queue.pop(0)
        if isinstance(item, list):
            for element in reversed(item):
                queue.insert(0, element)
        else:
            result.append(item)

    return result


if __name__ == "__main__":
    example = [1, 2, [3, 4, 5, [6, 7], 8, [9, 10], 11], [12, 13], 14, [15, 16], 17]
    result = flatten_queue(example)
    print(result)

    result = flatten_rec(example)
    print(result)
