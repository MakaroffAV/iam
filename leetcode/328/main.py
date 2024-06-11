# tags: linked list

from typing import Optional

class ListNode:
        def __init__(self, val=0, next=None):
            self.val  = val
            self.next = next

class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return head

        o = head
        e = eh = head.next
        
        while e and e.next:
        
            o.next = o.next.next
            o = o.next
            
            e.next = e.next.next
            e = e.next
            
        o.next = eh
        return head
