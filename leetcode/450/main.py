# tags: tree, binary search tree

from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val   = val
        self.left  = left
        self.right = right
        
class Solution:
    def min_node_val(self, node):
        c = node
        while c.left is not None:
            c = c.left
        return c
        
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if root is None:
            return root
        
        if key < root.val:
            root.left  = self.deleteNode(root.left,  key)
        elif key > root.val:
            root.right = self.deleteNode(root.right, key)
        else:
            if root.right is None:
                return root.left
            elif root.left is None:
                return root.right
        
            temp = self.min_node_val(root.right)
            root.val = temp.val
            root.right = self.deleteNode(root.right, temp.val)
        
        return root
    
if __name__ == "__main__":
    test_cases = [
        [
            [5,3,6,2,4,None,7],
            3,
            [5,4,6,2,None,None,7],
        ],
        [
            [5,3,6,2,4,None,7],
            0,
            [5,3,6,2,4,None,7],
        ],
        [
            [],
            0,
            [],
        ]
    ]
