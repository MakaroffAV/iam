# tags: dfs, graph

from typing import List

class Solution:
    def dfs(self, graph, node, visited):
        visited.add(node)
        for i in graph[node]:
            if i not in visited:
                self.dfs(graph, i, visited)
        
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        graph = {}
        for i in range(len(rooms)):
            graph[i] = []
            for j in rooms[i]:
                graph[i].append(j)

        visited = set()
        self.dfs(graph, 0, visited)
        return len(visited) == len(graph.keys())
    
if __name__ == "__main__":
    test_cases = [
        [
            [[1],[2],[3],[]],
            True
        ],
        [
            [[1,3],[3,0,1],[2],[0]],
            False
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().canVisitAllRooms(i[0]))
        