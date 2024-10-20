"""
Уставший профессор вечером играет в увлекательную игру.
Изначально на доске слева направо записаны целые числа 𝑎1,𝑎2,…,𝑎𝑛.
Пока не уснет, профессор каждую секунду смотрит на числа, видит, что самое левое равно 𝑥, а самое правое равно 𝑦.
Если 𝑥 меньше, то профессор радуется, стирает слева 𝑥, а справа дописывает (𝑥+𝑦)mod230.
Иначе профессор очень расстаивается, стирает 𝑦, а слева дописывает (𝑦−𝑥)mod230.
Студенты подсчитали, что перед сном профессор успел сделать ровно 𝑘 операций. Что было написано на доске, когда он
наконец заснул? Для простоты можно считать, что доска в обе стороны бесконечна.
"""
from collections import deque


def simulate_professor_game(arr: list[int], k: int) -> list[int]:
    mod_value = 2 ** 30
    dq = deque(arr)

    for i in range(k):
        left = dq[0]
        right = dq[-1]

        if left < right:
            dq.popleft()
            dq.append((left + right) % mod_value)
        else:
            dq.pop()
            dq.appendleft((right - left) % mod_value)

    return list(dq)

def read_data() -> tuple[int, list[int]]:
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    return k, arr


if __name__ == '__main__':
    k, arr = read_data()
    result = simulate_professor_game(arr, k)
    print(" ".join(map(str, result)))
