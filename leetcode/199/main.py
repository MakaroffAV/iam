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
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if not root:
            return res
        
        queue = deque([(root, 0)])
        
        while queue:
            size = len(queue)
            for _ in range(size):
                node, level = queue.popleft()
                
                if len(res)-1 < level:
                    res.append([])
                res[level].append(node.val)

                if node.left:
                    queue.append((node.left,  level + 1))
                if node.right:
                    queue.append((node.right, level + 1))

        return [i[-1] for i in res]
            
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
            [1,2,3,None,5,None,4],
            [1,3,4],
        ],
        [
            [1,None,3],
            [1,3],
        ],
        [
            [],
            []
        ],
    ]
    for i in test_cases:
        print(i[1] == Solution().rightSideView(create_tree(i[0])))
