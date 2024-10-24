"""
Дано натуральное число N. Рассмотрим его разбиение на натуральные слагаемые. Два разбиения, отличающихся только
порядком слагаемых, будем считать за одно, поэтому можно считать, что слагаемые в разбиении упорядочены по невозрастанию.
"""
def find_partitions(n, max_value=None, partition=None, partitions=None):
    if max_value is None:
        max_value = n
    if partition is None:
        partition = []
    if partitions is None:
        partitions = []

    if n == 0:
        partitions.append(partition)
        return [] if partitions == [[]] else None

    for i in range(min(n, max_value), 0, -1):
        find_partitions(n - i, i, partition + [i], partitions)

    return partitions


if __name__ == '__main__':
    number = int(input())
    result = find_partitions(number)
    result.sort(key=lambda item: -len(item))
    for partition in result:
        partition.sort(reverse=True)
        print(" ".join(map(str, partition)) + " ")
