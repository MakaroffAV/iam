# tags: binary_tree, dfs

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val   = val
        self.left  = left
        self.right = right
        
class Solution:
    def update(self):
        self.r += 1
    
    def inorder(self, node, max):
        if node is None:
            return
        
        if node.val >= max:
            self.update( )
            max = node.val
    
        self.inorder(node.left,  max)
        self.inorder(node.right, max)

    def goodNodes(self, root: TreeNode) -> int:
        self.r = 0
        if root is None:
            return self.r
        
        self.inorder(root, root.val)
        
        print(self.r)
        return self.r
        
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
            [3,1,4,3,None,1,5],
            4,
        ],
        [
            [3,3,None,4,2],
            3,
        ],
    ]
    for i in test_cases:
        print(i[1] == Solution().goodNodes(create_tree(i[0])))
