from typing import List, Tuple

import matplotlib.pyplot as plt
import numpy as np

AMOUNT_OF_EXPERIMENTS = 1000


def generate_random_sequence() -> np.array:
    length = np.random.randint(1, 100000)
    return np.random.binomial(length, 0.5, size=length)


def count_unique(int_list: np.array) -> int:
    return len(np.unique(int_list))


def display_plot(x_values: List[int], y_values: List[int]) -> None:
    plt.plot(x_values, y_values)
    plt.xlabel("Sequence length")
    plt.ylabel("Unique numbers count")
    plt.title("Unique numbers")
    plt.show()


if __name__ == "__main__":
    results: List[Tuple[int, int]] = []

    for _ in range(AMOUNT_OF_EXPERIMENTS):
        random_sequence = generate_random_sequence()
        sequence_length = len(random_sequence)
        unique_numbers_count = count_unique(random_sequence)
        results.append((sequence_length, unique_numbers_count))

    results.sort(key=lambda value: value[0])
    x, y = zip(*results)

    display_plot(x, y)

