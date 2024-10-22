import random


def quick_select(elements: list[int], n:int) -> int:
    if len(elements) == 1:
        return elements[0]

    left, middle, right = partition(elements, random.choice(elements))
    if len(left) > n:
        return quick_select(left, n)
    elif len(left) + len(middle) < n:
        return quick_select(right, n - len(left) - len(middle))

    return middle[0]

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
