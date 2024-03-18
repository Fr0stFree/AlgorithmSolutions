"""
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.
"""
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left_index = 0
        right_index = len(nums) - 1

        if target <= nums[0]:
            return 0

        if target > nums[-1]:
            return len(nums)

        while right_index - left_index != 1:
            middle_index = (left_index + right_index) // 2
            middle_value = nums[middle_index]
            if middle_value == target:
                return middle_index

            if middle_value < target:
                left_index = middle_index

            elif middle_value > target:
                right_index = middle_index

        return right_index


if __name__ == '__main__':
    solution = Solution()

    assert solution.searchInsert([1, 3, 5, 6, 7], 5) == 2
    assert solution.searchInsert([1, 3, 5, 14, 15], 7) == 3
    assert solution.searchInsert([1, 3, 5, 6], 2) == 1
    assert solution.searchInsert([1, 3, 5, 6], 7) == 4
