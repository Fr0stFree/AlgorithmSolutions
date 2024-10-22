"""
В очереди в магазин стоят люди. Человек 𝑖 хочет купить товар 𝑎𝑖. Изначально в магазине ничего нет.
Происходят события следующих типов:
    В момент времени 𝑇 поступил один экземпляр товара 𝐴.
    В момент времени 𝑇 в конец очереди встал человек, который хочет купить товар 𝐴.
Нужно промоделировать процесс и для каждого человека определить, сколько он будет стоять в очереди.
Замечание: как только первый в очереди может купить то, что хочет, он сразу мгновенно покупает и уходит.
"""
from collections import defaultdict, deque

def solve(events: list[tuple[int, int, int]]):
    inventory: dict[int, int] = defaultdict(lambda: 0)
    queue: deque[tuple[int, int]] = deque()
    result: list[int] = []

    for event in events:
        if event[0] == 1:
            inventory[event[2]] += 1

        else:
            queue.append((event[1], event[2]))

        while queue and queue[0][1] in inventory:
            customer = queue.popleft()
            inventory[customer[1]] -= 1
            result.append(event[1] - customer[0])

            if inventory[customer[1]] <= 0:
                inventory.pop(customer[1])

    result.extend([-1] * len(queue))
    return result


def read_data() -> list[tuple[int, int, int]]:
    size = int(input())
    return [tuple(map(int, input().split(" "))) for _ in range(size)]


if __name__ == '__main__':
    data = read_data()
    solution = solve(data)
    print(" ".join(map(str, solution)))
