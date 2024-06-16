# tags: linked_list, two_pointers, stack

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
    def pairSum(self, head: Optional[ListNode]) -> int:
        
        if not head:
            return 0
        
        n = 0
        s = []
        c = head
        while c is not None:
            n += 1
            s.append(c.val)

            c = c.next
            
        m = 0
        
        l = 0
        r = len(s)-1

        while l < r:
            t = s[l] + s[r]
            l += 1
            r -= 1
            if t > m:
                m = t
        return m   

if __name__ == "__main__":
    test_cases = [
        [
            [5,4,2,1],
            6,
        ],
        [
            [4,2,2,3],
            7,
        ],
        [
            [1,100000],
            100001,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().pairSum(create_linked_list(i[0])))
