"""
8б класс решил на слет взять много Спрайта. Для этого они собрались сконструировать переносной холодильник 𝑎×𝑏×𝑐,
который будет вмещать ровно 𝑛 кубических банок Спрайта размером 1×1×1. Чтобы лимонад доехал как можно более холодным,
они хотят минимизировать теплопотери; то есть минимизировать площадь поверхности.

Например, если емкость холодильника должна равняться 12, то возможны следующие варианты:
    3 2 2 → 32
    4 3 1 → 38
    6 2 1 → 40
    12 1 1 → 50
В этом примере оптимальным является холодильник 322. Помогите 8б найти оптимальный холодильник в общем случае.
"""
import math


def find_smallest_node_lengths(volume: int) -> list[int]:
    smallest_area = math.inf
    best_cube = [0, 0, 0]

    for a in range(1, int(math.isqrt(volume)) + 1):
        if volume % a != 0:
            continue

        for b in range(a, int(math.isqrt(volume // a)) + 1):
            if (volume // a) % b != 0:
                continue

            c = volume // (a * b)
            area = 2 * (a * b + b * c + c * a)

            if area < smallest_area:
                smallest_area = area
                best_cube = [a, b, c]

    return best_cube


if __name__ == '__main__':
    volume = int(input())
    result = find_smallest_node_lengths(volume)
    print(" ".join(map(str, sorted(result))))
