from typing import Sequence


def is_sum_exist(expected_sum: int, numbers: Sequence[int]) -> bool:
    
    def find_summands(total: int) -> tuple[bool, list[int]]:
        if total == 0:
            return True, []
        if total < 0:
            return False, []

        for number in numbers:
            new_total = total - number
            is_exist, summands = find_summands(new_total)
            if is_exist:
                summands.append(number)
                return True, summands
        return False, []

    is_exist, summands = find_summands(expected_sum)
    print(summands)
    return is_exist


if __name__ == "__main__":
    numbers = [2, 4, 5, 7]
    result = is_sum_exist(9, numbers)
    print(result)
