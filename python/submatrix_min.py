"""
Дана матрица 𝑛×𝑛, состоящая из целых чисел. Для каждой её подматрицы размера 𝐿×𝐿 найдите минимум в этой подматрице.
Подматрицей здесь называется "подпрямоугольник".
Внимание. Решение должно работать за 𝑂(𝑛2).
"""
def solve(matrix: list[list[int]], window_size: int) -> list[list[int]]:
    new = []
    for row, line in enumerate(matrix):
        new.append([])
        for index in range(0, len(line)-window_size+1):
            res = min(line[index:index+window_size])
            new[row].append(res)

    result = []
    for row in range(len(new)-window_size+1):
        result.append([])
        for column in range(len(new[0])):
            res = min([new[row+i][column] for i in range(window_size)])
            result[row].append(res)

    return result


def read_data() -> tuple[list[list[int]], int]:
    size, window = map(int, input().split(" "))
    matrix = [list(map(int, input().split(" "))) for _ in range(size)]
    return matrix, window


if __name__ == "__main__":
    data = read_data()
    result = solve(*data)
    print("\n".join(" ".join(map(str, line)) for line in result))
