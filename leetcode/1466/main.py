# tags: graph, dfs

from typing import List

class Solution:
    
    def __init__(self):
        self.rotations = 0
        
    def minReorder(self, n: int, connections: List[List[int]]) -> int:
        
        def dfs(node, target, visited, edges):
            visited.add(node)
            if node == target:
                return
            for i in graph[node]:
                if i not in visited:
                    if edges[(node, i)]:
                        self.rotations += 1
                        edges[(i, node)] = True
                        edges[(node, i)] = False

                    dfs(i, target, visited, edges)
        
        edges = {}
        graph = {}
        nodes = set()
        
        for i in connections:
            
            edges[(i[0], i[1])] = True
            edges[(i[1], i[0])] = False
                
            if not graph.get(i[0]):
                graph[i[0]] = [i[1]]
            else:
                graph[i[0]].append(i[1])
            
            if not graph.get(i[1]):
                graph[i[1]] = [i[0]]
            else:
                graph[i[1]].append(i[0])
                
            if i[0] not in nodes and i[0] != 0:
                nodes.add(i[0])
            if i[1] not in nodes and i[1] != 0:
                nodes.add(i[1])
        
        visited = set()
        dfs(0, i, visited, edges)

        return self.rotations
    
if __name__ == "__main__":
    test_cases = [
        [
            6,
            [[0,1],[1,3],[2,3],[4,0],[4,5]],
            3
        ],
        [
            5,
            [[1,0],[1,2],[3,2],[3,4]],
            2,
        ],
        [
            3,
            [[1,0],[2,0]],
            0
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().minReorder(i[0], i[1]))