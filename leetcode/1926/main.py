# tags: bfs, maze

from typing import List

class Solution:
    
    def __init__(self) -> None:
        self.row_d = [-1, 0, 1, 0]
        self.col_d = [0, 1, 0, -1]
        
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        
        def is_valid(row, col):
            if row < 0 or col < 0 or row >= rows or col >= cols:
                return False
            if v[row][col]:
                return False
            if maze[row][col] == "+":
                return False
            return True
        
        q = []
        
        rows = len(maze)
        cols = len(maze[0])
        
        min_path = float("inf")
        
        q.append(entrance+[1])
        v = [
            [
                False 
                for _ in range(cols)
            ]
            for  _   in  range(rows)
        ]
        v[entrance[0]][entrance[1]] = True
        
        while len(q) > 0:
            x = q[0][0]
            y = q[0][1]   
            s = q[0][2]     

            q = q[1:]
            
            for i in range(4):
                adj_x = x + self.row_d[i]
                adj_y = y + self.col_d[i]
                
                if is_valid(adj_x, adj_y):
                    if maze[adj_x][adj_y] == "." and ((adj_x == 0 or adj_x == rows-1) or (adj_y == 0 or adj_y == cols-1)):
                        if s < min_path:
                            min_path = s

                    v[adj_x][adj_y] = True
                    q.append((adj_x, adj_y, s+1))
                    
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
        ],
        [
            [
                ["+",".","+","+","+","+","+"],
                ["+",".","+",".",".",".","+"],
                ["+",".","+",".","+",".","+"],
                ["+",".",".",".",".",".","+"],
                ["+","+","+","+",".","+","."],
            ],
            [
                0,
                1,
            ],
            7
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().nearestExit(i[0], i[1]))
