# tags: dfs, graph

from typing import List

class Solution:
    
    def all_visited(self, visited):
        for i in visited:
            if visited[i] == False:
                return False
        return True

    def dfs(self, graph, node, visited):
        visited[node]  = True
        for i in  graph[node]:
            if not visited[i]:
                self.dfs(graph, i, visited)

    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        graph =  {}
        c = isConnected
        for i, v1 in enumerate(c):
            for j, v2 in enumerate(v1):
                if i == j:
                    pass
                else:
                    if graph.get(i+1) is None:
                        graph[i+1] = []
                    if v2 == 1:
                        graph[i+1].append(j+1)
                        
        visited = {i: False for i in graph}
        
        regions = 0
        while not self.all_visited(visited):
            for i in visited:
                if visited[i] == False:
                    regions += 1
                    self.dfs(graph, i, visited)
        return regions
    
if __name__ == "__main__":
    test_cases = [
        [
            [[1,1,0],[1,1,0],[0,0,1]],
            2,
        ],
        [
            [[1,0,0],[0,1,0],[0,0,1]],
            3
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().findCircleNum(i[0]))
        