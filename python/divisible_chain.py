import math


def divisible_chain(start: int, divisor: int) -> str:

    def get_divisors(n: int) -> tuple[int, list[int]]:
        if n < 1:
            return int(math.inf), []

        if n == 1:
            return 1, [n]

        score1, path1 = get_divisors(n - 1)
        score1 += 1

        if n % divisor == 0:
            score2, path2 = get_divisors(n // divisor)
            score2 += 1

            if score2 < score1:
                path2.append(n)
                return score2, path2

        path1.append(n)
        return score1, path1

    _, path = get_divisors(start)
    return " -> ".join(map(str, reversed(path)))


if __name__ == "__main__":
    print(divisible_chain(100, 3))
    assert divisible_chain(29, 3) == "29 -> 28 -> 27 -> 9 -> 3 -> 1"
