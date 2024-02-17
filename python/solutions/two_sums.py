"""
Given an array of integers nums and an integer target, return indices of the two numbers such that
they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
"""
from collections import defaultdict
from typing import List, Dict


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        left_index, right_index = 0, len(nums) - 1
        extra_dict: Dict[int, List[int]] = defaultdict(list)
        
        [extra_dict[value].append(index) for index, value in enumerate(nums)]
        nums.sort()
        
        while left_index <= right_index:
            elems_sum = nums[left_index] + nums[right_index]
            if elems_sum < target:
                left_index += 1
                continue
                
            if elems_sum > target:
                right_index -= 1
                continue
                
            return [extra_dict[nums[left_index]].pop(), extra_dict[nums[right_index]].pop()]
            
            
if __name__ == "__main__":
    solution = Solution()
    
    assert sorted(solution.twoSum([2, 7, 11, 15], 9)) == [0, 1]
    assert sorted(solution.twoSum([3, 2, 4], 6)) == [1, 2]
    assert sorted(solution.twoSum([3, 3, 1], 6)) == [0, 1]
