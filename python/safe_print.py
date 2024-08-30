import time

from threading import Thread, Lock
from queue import Queue, Full
from typing import Any

queue = Queue(maxsize=4)
electronics = ["смартфон", "ноутбук", "планшет", "камера", "гарнитура",
               "телевизор", "гаджет", "монитор", "роутер", "плеер"]

def get_safe_printer() -> callable:
    lock = Lock()

    def _print(*args, **kwargs) -> None:
        with lock:
            print(*args, **kwargs)

    return _print

safe_print = get_safe_printer()

def consume(queue: Queue) -> None:
    while not queue.empty():
        element = queue.get()
        time.sleep(1)
        safe_print(f"Продан товар: {element}")

def produce(queue: Queue, elements: Any) -> None:
    for element in elements:
        try:
            queue.put_nowait(element)
            safe_print(f"Поставлен товар: {element}")
        except Full:
            safe_print(f"Склад временно заполнен")
            elements.append(element)
        finally:
            time.sleep(.5)

    safe_print("Поставки закончились")

producer = Thread(target=produce, args=(queue, electronics), daemon=True)
producer.start()
consumer = Thread(target=consume, args=(queue,), daemon=True)
consumer.start()

consumer.join()
