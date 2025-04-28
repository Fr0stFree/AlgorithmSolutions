from collections import defaultdict


def read_data() -> list[str]:
    amount = int(input())
    words = [input() for _ in range(amount)]
    return words


def count_words(words: list[str]) -> tuple[int, list[int]]:
    visited = defaultdict(lambda: 0)
    for word in words:
        visited[word] += 1
    return len(visited), list(visited.values())


if __name__ == "__main__":
    words = read_data()
    amount, counts = count_words(words)
    print(f"{amount}\n{' '.join(map(str, counts))}")
