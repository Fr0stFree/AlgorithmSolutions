def selection_sort(items: list[int]) -> list[int]:
    for index, value in enumerate(items):
        min_item: tuple[int, int] = (index, value)
        for next_index, next_value in enumerate(items[index:]):
            if next_value < min_item[1]:
                min_item = next_index + index, next_value

        items[index], items[min_item[0]] = min_item[1], items[index]
    return items
