# tags: binary_tree, dfs

from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val   = val
        self.left  = left
        self.right = right
        
class Solution:
    def update(self, v):
        if v > self.m:
            self.m = v

    def inorder(self, node, l=0):
        if node is None:
            return None
        self.inorder(node.left,  l+1)
        self.update(l)
        self.inorder(node.right, l+1)
        
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        self.m = 0
        if root is None:
            return self.m
        self.inorder(root, 1)
        
        return self.m

def create_tree(a: list):
    if len(a) == 0:
        return None
    if len(a) == 1:
        return TreeNode(a[0])
    
    def inner(index=0):
        if len(a)-1 < index:
            return None
        if a[index] is None:
            return None
        
        node = TreeNode(a[index])
        node.left  = inner(index * 2 + 1)
        node.right = inner(index * 2 + 2)
        
        return node
    
    return inner()
    
if __name__ == "__main__":
    test_cases = [
        [
            [3,9,20,None,None,15,7],
            3,
        ],
        [
            [1,None,2],
            2,
        ],
        [
            [1,2],
            2
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().maxDepth(create_tree(i[0])))
