# tags: graph, dfs

from typing import List

class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        graph = {}
        
        def dfs(visited, node, target, path=1):
            visited.add(node)
            
            if node == target:
                return path
            
            for i in graph[node]:
                if i[0] not in visited:
                    res = dfs(visited, i[0], target, path*i[1])
                    if res is not None and res != -1:
                        return res
            return -1
        
        def find(src, dst):
            if src not in graph.keys() or dst not in graph.keys():
                return -1
            
            if src == dst:
                return  1
               
            visited = set()        
            return dfs(visited, src, dst, 1)
        
        def update_graph(src, dst, v):
            if graph.get(src) is None:
                graph[src] = [(dst, v)]
            else:
                graph[src].append((dst, v))
                
        for i, v in enumerate(equations):
            update_graph(v[0], v[1], values[i])
            update_graph(v[1], v[0], 1/values[i])
            
        res = []
        for i in queries:
            res.append(find(i[0], i[1]))
            
        print(res)
        return res

if __name__ == "__main__":
    test_cases = [
        [
            [["a","b"],["b","c"]],
            [2.0,3.0],
            [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]],
            [6.00000,0.50000,-1.00000,1.00000,-1.00000]
        ],
        [
            [["x1","x2"],["x2","x3"],["x3","x4"],["x4","x5"]],
            [3.0,4.0,5.0,6.0],
            [["x1","x5"],["x5","x2"],["x2","x4"],["x2","x2"],["x2","x9"],["x9","x9"]],
            None
        ]
    ]
    for i in test_cases:
        print(i[3] == Solution().calcEquation(i[0], i[1], i[2]))
        