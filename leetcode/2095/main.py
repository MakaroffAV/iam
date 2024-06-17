# tags: linked_list, two_pointers

from typing import List
from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val  = val
        self.next = next
        
def create_linked_list(a: List) -> ListNode:
    h = ListNode(a[0])
    c = h
    for i in a[1:]:
        c.next = ListNode(i)
        c = c.next
    return h

class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None
        if head.next is None:
            return []
        n = 0
        c = head
        while c is not None:
            n += 1
            c = c.next
        m = (n // 2) if n % 2 == 0 else (n // 2)
        
        n = 0
        c = head
        while c is not None:
            if n == m-1:
                c.next = c.next.next if c.next.next is not None else None
                break
            n += 1
            c = c.next
        return head
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,3,4,7,1,2,6],
            [1,3,4,1,2,6],
        ],
        [
            [1,2,3,4],
            [1,2,4],
        ],
        [
            [2,1],
            [2]
        ],
    ]
    for i in test_cases:
        print(create_linked_list(i[1]) == Solution().deleteMiddle(create_linked_list(i[0])))
