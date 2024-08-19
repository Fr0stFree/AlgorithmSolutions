"""
Напишите класс FilterQueue со следующими свойствами:
Это потомок asyncio.Queue
В экземпляре класса атрибут очередь.window содержит первый элемент очереди, или None, если очередь пуста (просмотр очередь.window не влияет на состояние очереди)
С помощью операции фильтр in очередь можно определить, присутствуют ли в очереди такие элементы, что выражение фильтр(элемент) истинно
Метод .later() синхронно переставляет первый элемент очереди в её конец, или вызывает исключение asyncio.QueueEmpty, если очередь пуста
Метод .get() содержит необязательный параметр фильтр. Вызов очередь.get(фильтр) работает так:
Если в очереди нет элементов, на которых фильтр(элемент) истинно, работает как обычный .get().
Если в очереди есть элементы, на которых фильтр(элемент) истинно, переставляет первый элемент очереди в её конец до тех пор, пока фильтр(элемент) не истинно, а затем выполняет обычный .get().
"""
import asyncio
from typing import TypeVar, Callable

import pytest

T = TypeVar('T')
Predicate = Callable[[T], bool]


class FilterQueue(asyncio.Queue):
    def window(self) -> T | None:
        if self.empty():
            return None
        return self._queue[0]

    def later(self) -> None:
        self.put_nowait(self.get_nowait())

    def __contains__(self, predicate: Predicate) -> bool:
        return any(predicate(item) for item in self._queue)

    async def get(self, predicate: Predicate = None) -> T:
        if predicate is None or predicate not in self:
            return await super().get()

        while not predicate(item := await super().get()):
            await self.put(item)

        return item


# Tests

def test_window():
    queue = FilterQueue()
    queue.put_nowait(1)

    assert queue.window() == 1

    queue.put_nowait(2)

    assert queue.window() == 1

    queue.get_nowait()
    queue.get_nowait()

    assert queue.window() is None


def test_later():
    queue = FilterQueue()

    with pytest.raises(asyncio.QueueEmpty):
        queue.later()

    for number in (1, 2, 3, 4, 5):
        queue.put_nowait(number)

    assert queue.window() == 1

    queue.later()

    assert queue.window() == 2


def test_contains():
    queue = FilterQueue()
    for number in (1, 2, 3, 4, 5):
        queue.put_nowait(number)

    predicate = lambda element: element % 2 == 0
    assert predicate in queue

    predicate = lambda element: element == 'Protsanov'
    assert predicate not in queue


@pytest.mark.asyncio
async def test_get_without_predicate():
    queue = FilterQueue()
    for number in (1, 2):
        queue.put_nowait(number)

    assert await queue.get() == 1
    assert await queue.get() == 2
    with pytest.raises(asyncio.QueueEmpty):
        queue.get_nowait()


@pytest.mark.asyncio
async def test_get_with_falsy_predicate():
    queue = FilterQueue()
    predicate = lambda element: element == 'Protsanov'
    for number in (1, 2):
        queue.put_nowait(number)

    assert await queue.get(predicate) == 1
    assert await queue.get(predicate) == 2
    with pytest.raises(asyncio.QueueEmpty):
        queue.get_nowait()


@pytest.mark.asyncio
async def test_get_with_truthy_predicate():
    queue = FilterQueue()
    predicate = lambda element: element == 'Protsanov'
    for number in (1, 'Protsanov', 'Protsanov', 2):
        queue.put_nowait(number)

    assert await queue.get(predicate) == 'Protsanov'
    assert await queue.get(predicate) == 'Protsanov'
    assert await queue.get(predicate) == 2
    assert await queue.get(predicate) == 1
    with pytest.raises(asyncio.QueueEmpty):
        queue.get_nowait()
