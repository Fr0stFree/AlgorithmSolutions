from functools import partial
from typing import Callable


def chain_sum(number: int | None = None, accumulated: int = 0) -> Callable | int:
    if number is None:
        return accumulated

    accumulated += number
    return partial(chain_sum, accumulated=accumulated)


assert chain_sum(50)(40)(30)() == 120
assert chain_sum() == 0
