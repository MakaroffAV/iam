# tags: binary_tree, dfs

from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val   = val
        self.left  = left
        self.right = right
        
class Solution:
    def traversal(self, node, res):
        if node.left is None and node.right is None:
            res.append(node.val)
        
        if node.left is not None:
            self.traversal(node.left,  res)
        if node.right is not None:
            self.traversal(node.right, res)
    
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        if root1 is None and root2 is None:
            return True
        if root1 is None and root2 is not None:
            return False
        if root1 is not None and root2 is None:
            return False
        
        res1 = []
        self.traversal(root1, res1)
        
        res2 = []
        self.traversal(root2, res2)
        
        return res1 == res2
        

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
            [3,5,1,6,2,9,8,None,None,7,4],
            [3,5,1,6,7,4,2,None,None,None,None,None,None,9,8],
            True,
        ],
        [
            [1,2,3],
            [1,3,2],
            False,
        ],
    ]
    for i in test_cases:
        print(i[2] == Solution().leafSimilar(create_tree(i[0]), create_tree(i[1])))
