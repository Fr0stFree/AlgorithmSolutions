def main(n: int, k: int) -> None:
    is_divisible = lambda divisible, divider: divisible % divider == 0
    result = [n]

    while n != 1:
        if is_divisible(n, k):
            n //= k
        else:
            n -= 1
        result.append(n)

    print("->".join(map(str, result)))
    return None


if __name__ == '__main__':
    main(29, 3)