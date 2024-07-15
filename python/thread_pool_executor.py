from concurrent.futures import ThreadPoolExecutor, Future
import threading
import time

strings = [
    "Да Здравствует ThreadPoolExecutor!!!",
    "Многопоточность в Python позволяет выполнять несколько задач одновременно, улучшая производительность.",
    "Многопоточность может увеличить сложность управления памятью и ресурсами.",
    "Правильное использование многопоточности в Python может значительно улучшить производительность приложений."
]


def to_uppercase(string: str) -> str:
    return string.upper()


def main():
    result = []

    with ThreadPoolExecutor(max_workers=3) as executor:
        futures = []
        for string in strings:
            futures.append(executor.submit(to_uppercase, string))

        for future in futures:
            result.append(future.result())

    [*map(print, result)]


if __name__ == '__main__':
    main()
