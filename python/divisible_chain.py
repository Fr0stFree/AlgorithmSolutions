def divisible_chain(n: int, k: int) -> str:
    is_divisible = lambda divisible, divider: divisible % divider == 0
    result = [n]

    while n != 1:
        if is_divisible(n, k):
            n //= k
        else:
            n -= 1
        result.append(n)

    return " -> ".join(map(str, result))


if __name__ == '__main__':
    assert divisible_chain(29, 3) == '29 -> 28 -> 27 -> 9 -> 3 -> 1'
