from functools import partial
from typing import Callable


def chain_sum(number: int | None = None, accumulated: int = 0) -> Callable | int:
    if number is None:
        return accumulated

    accumulated += number
    return partial(chain_sum, accumulated=accumulated)


def chain_sum_2(result: int | None = None) -> Callable | int:

    def inner(number: int | None = None):
        nonlocal result
        if number is None:
            return result or 0

        result += number
        return chain_sum_2(result)

    return inner() if result is None else inner


assert chain_sum(50)(40)(30)() == 120
assert chain_sum() == 0
assert chain_sum_2(50)(40)(30)() == 120
assert chain_sum_2() == 0
