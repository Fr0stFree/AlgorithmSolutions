from contextlib import suppress
from math import inf, isinf


class BinHeap[T]:
    def __init__(self) -> None:
        self._collection = [None]

    def add(self, elem: T) -> None:
        self._collection.append(elem)
        self._sift_up(index=len(self._collection)-1)

    def pop(self) -> T:
        if len(self._collection) == 1:
            raise ValueError("Heap is empty")

        if len(self._collection) == 2:
            return self._collection.pop()

        elem = self._collection[1]
        self._collection[1] = self._collection.pop()
        self._sift_down(index=1)
        return elem

    def _sift_down(self, index: int) -> None:
        left, right = inf, inf
        with suppress(IndexError):
            left = self._collection[index * 2]
        with suppress(IndexError):
            right = self._collection[index * 2 + 1]

        elem = self._collection[index]
        minimum = min(left, right, elem)
        if minimum == elem:
            return
        if minimum == left:
            self._collection[index], self._collection[index*2] = self._collection[index*2], elem
            self._sift_down(index=index*2)
        if minimum == right:
            self._collection[index], self._collection[index*2+1] = self._collection[index*2+1], elem
            self._sift_down(index=index*2+1)

    def _sift_up(self, index: int) -> None:
        elem = self._collection[index]
        ancestor = self._collection[index // 2]
        if ancestor is None:  # the elem become the root
            return
        if ancestor > elem:
            self._collection[index // 2], self._collection[index] = self._collection[index], self._collection[index // 2]
            return self._sift_up(index//2)

heap = BinHeap()
elems = [1, 5, 3, 8, 0, 4]
for elem in elems:
    heap.add(elem)

for _ in range(len(elems)):
    elem = heap.pop()
    print(elem)