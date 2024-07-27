import threading
import time
from collections import Counter, defaultdict
from concurrent.futures import ThreadPoolExecutor

messages = [
    "Привет, коллеги! Кто-нибудь может объяснить разницу между асинхронным и многопоточным программированием в Python?",
    "Асинхронное программирование позволяет обрабатывать несколько задач без блокировки основного потока.",
    "Многопоточное программирование также обеспечивает параллельное выполнение, но с использованием потоков вместо асинхронных задач.",
    "Помните, что асинхронное программирование может требовать пересмотра структуры кода и логики приложения.",
    "Многопроцессорное программирование может использовать модуль multiprocessing.Event для организации событийного взаимодействия между процессами.",
    "Асинхронные функции могут использовать asyncio.ensure_future() для запуска задачи в асинхронном режиме."
]
result = defaultdict(lambda: 0)

lock = threading.Lock()


def count_letters(sentence: str) -> None:
    letters = Counter(filter(lambda letter: letter not in ['.', ',', '!', '?', '-', ' ', '_', '(', ')', "'", '"', '/'],
                             sentence.lower()))
    with lock:
        for letter, amount in letters.items():
            result[letter] += amount


with ThreadPoolExecutor() as executor:
    executor.map(count_letters, messages)
    time.sleep(0.1)
    sorted_letters = dict(sorted(result.items(), key=lambda item: item[0]))
    print(sorted_letters)
