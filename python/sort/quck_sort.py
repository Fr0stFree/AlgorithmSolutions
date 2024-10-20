import random

def quick_sort(elements: list[int]) -> list[int]:
    if len(elements) <= 1:
        return elements

    left, middle, right = partition(elements, random.choice(elements))

    return quick_sort(left) + middle + quick_sort(right)

def partition(elements: list[int], pivot: int) -> tuple[list[int], list[int], list[int]]:
    left, middle, right = [], [], []
    for element in elements:
        if element < pivot:
            left.append(element)
        elif element > pivot:
            right.append(element)
        else:
            middle.append(element)
    return left, middle, right
