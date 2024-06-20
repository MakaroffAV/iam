# tags: linked_list

from typing import List
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val  = val
        self.next = next
        
def create_linked_list(a: List) -> ListNode:
    if not a:
        return None
    h = ListNode(a[0])
    c = h
    for i in a[1:]:
        c.next = ListNode(i)
        c = c.next
    return h

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None
        prev = None
        curr = head
        while curr is not None:
            next = curr.next
            curr.next = prev
            
            prev = curr
            curr = next
        head = prev
        return head
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,3,4,5],
            [5,4,3,2,1],
        ],
        [
            [1,2],
            [2,1],
        ],
        [
            [],
            []
        ]
    ]
    for i in test_cases:
        print(create_linked_list(i[1]) == Solution().reverseList(create_linked_list(i[0])))
