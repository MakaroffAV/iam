# tags: bfs

from typing import List

class Solution:
    
    def __init__(self) -> None:
        self.row_d = [-1, 0, 1, 0]
        self.col_d = [0, 1, 0, -1]
        
    def is_valid(self, visited, row, col):
        if row < 0 or col < 0 or row >= self.rows or col >= self.cols:
            return False
        
        if visited[row][col]:
            return False
        
        if self.maze[row][col] == "+":
            return False
        
        return True
        
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        q = []
        v = {}
        
        self.maze = maze
        self.rows = len(maze)
        self.cols = len(maze[0])
        
        min_path = float("inf")
        cur_step = 0
        
        q.append(entrance)
        v = [
            [
                False 
                for i in range(self.cols)
            ]
            for i in range(self.rows)
        ]
        e = [
            [
                self.maze[i2][i1] == "." and ((i2 == 0 or i2 == self.rows-1) or (i1 == 0 or i1 == self.cols-1))
                for i1 in range(self.cols)
            ] 
            for i2 in range(self.rows)
        ]
        
        while len(q) > 0:            
            x, y = q[0]
            q    = q[1:]
            
            cur_step += 1
            
            for i in range(4):
                adj_x = x + self.row_d[i]
                adj_y = y + self.col_d[i]
                
                if self.is_valid(v, adj_x, adj_y):
                    if e[adj_x][adj_y]:
                        dist = abs(adj_x - entrance[0]) + abs(adj_y - entrance[1])
                        if dist != 0:
                            if dist < min_path:
                                min_path = dist
                                print(dist, adj_x, adj_y, cur_step)
                    v[adj_x][adj_y] = True
                    q.append((adj_x, adj_y))
                    
        return -1 if min_path == float("inf") else min_path
    
if __name__ == "__main__":
    test_cases = [
        [
            [
                ["+","+",".","+"],
                [".",".",".","+"],
                ["+","+","+","."],
            ],
            [
                1,
                2,
            ],
            1
        ],
        [
            [
                ["+","+","+"],
                [".",".","."],
                ["+","+","+"],
            ],
            [
                1,
                0,
            ],
            2
        ],
        [
            [
                [".","+"]
            ],
            [
                0,
                0,
            ],
            -1
        ],
        [
            [
                ["+",".","+","+","+","+","+"],
                ["+",".","+",".",".",".","+"],
                ["+",".","+",".","+",".","+"],
                ["+",".",".",".","+",".","+"],
                ["+","+","+","+","+",".","+"],
            ],
            [
                0,
                1,
            ],
            12
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().nearestExit(i[0], i[1]))
