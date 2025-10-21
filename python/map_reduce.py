from collections import defaultdict
import sys
from typing import Any, Generator, Protocol, TypeVar

TKey = TypeVar("TKey")
TValue = TypeVar("TValue")


class IMapper(Protocol):
    def map(self, data: str) -> Generator[tuple[TKey, TValue], None, None]: ...


class IReducer(Protocol):
    def reduce(self, key: TKey, values: list[TValue]) -> Any: ...


class MapReduce:
    def __init__(self, mapper: IMapper, reducer: IReducer):
        self.mapper = mapper
        self.reducer = reducer

    def shuffle(self, items: list[tuple[TKey, TValue]]) -> dict[TKey, list[TValue]]:
        result = defaultdict(list)
        for key, value in items:
            result[key].append(value)
        for key in result:
            result[key].sort()
        return result

    def get_splits(self) -> Generator[str, None, None]:
        while line := sys.stdin.readline():
            yield line

    def execute(self):
        to_shuffle = []
        for split in self.get_splits():
            for key, value in self.mapper.map(split):
                to_shuffle.append((key, value))

        for key, values in self.shuffle(to_shuffle).items():
            self.reducer.reduce(key, values)


### Implementation Example ###
class Mapper:
    def map(self, data: str) -> Generator[tuple[TKey, TValue], None, None]:
        fields = data.strip().split("\t")
        key, marker = fields
        yield key, marker


class Reducer:
    def reduce(self, key: TKey, values: list[TValue]) -> Any:
        if len(values) == 1:
            print(key)


if __name__ == "__main__":
    mapper = Mapper()
    reducer = Reducer()
    mr = MapReduce(mapper, reducer)
    mr.execute()
