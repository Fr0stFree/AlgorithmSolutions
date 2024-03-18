"""
Given an integer array nums sorted in non-decreasing order,
remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
Consider the number of unique elements of nums be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the unique elements in the order they were present
in nums initially. The remaining elements of nums are not important as well as the size of nums. Return k.
"""
from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        storage = set()

        for index, number in enumerate(nums.copy()):
            if number not in storage:
                storage.add(number)
            else:
                nums.pop(len(storage) - 1)

        return len(storage)


if __name__ == '__main__':
    solution = Solution()
    input_data = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]

    assert solution.removeDuplicates(input_data) == 5
    assert input_data == [0, 1, 2, 3, 4]

    input_data = [1, 2, 2]
    assert solution.removeDuplicates(input_data) == 2
    assert input_data == [1, 2]
