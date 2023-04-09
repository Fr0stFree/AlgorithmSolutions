"""
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
"""
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        result = []
        remainder = None
        node = None

        while remainder is None:
            if not all([list1, list2]):
                remainder = list2 if list1 is None else list1
                while remainder:
                    result.append(remainder.val)
                    remainder = remainder.next
                break

            if list1.val < list2.val:
                result.append(list1.val)
                list1 = list1.next

            else:
                result.append(list2.val)
                list2 = list2.next
        
        result.reverse()
        for number in result:
            node = ListNode(val=number, next=node)
    
        return node
        
                
            
        
        
if __name__ == '__main__':
    pass