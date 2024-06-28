# tags: bfs, binary_tree

from typing import List
from typing import Optional

from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val   = val
        self.left  = left
        self.right = right

class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        lmax = {}
        if not root:
            return 0
        
        queue = deque([(root, 0)])
        
        while queue:
            size = len(queue)
            for _ in range(size):
                node, level = queue.popleft()
                
                if lmax.get(level) is None:
                    lmax[level] = node.val
                else:
                    lmax[level] += node.val
                
                if node.left:
                    queue.append([node.left,  level+1])
                if node.right:
                    queue.append([node.right, level+1])
        
        max_lvl = 0
        max_val = float("-inf") 
        for i in lmax:
            if lmax[i] > max_val:
                max_lvl = i+1
                max_val = lmax[i]
        return max_lvl
            
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
            [1,7,0,7,-8,None,None],
            2,
        ],
        [
            [989,None,10250,98693,-89388,None,None,None,-32127],
            2
        ],
    ]
    for i in test_cases:
        print(i[1] == Solution().maxLevelSum(create_tree(i[0])))
