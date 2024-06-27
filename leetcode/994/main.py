# tags: matrix, bfs

from typing import List

class Solution:
    def __init__(self):
        self.row_d = [0, 0, 1, -1]
        self.col_d = [1, -1, 0, 0]
        
    def orangesRotting(self, grid: List[List[int]]) -> int:
        
        time = 0
        oranges_fresh = 0
        
        def is_valid(row, col):
            if row < 0 or col < 0 or row >= rows or col >= cols:
                return False
            if v[row][col]:
                return False
            if grid[row][col] == 0:
                return False
            if grid[row][col] == 2:
                return False
            return True
            
        rows = len(grid)
        cols = len(grid[0])
        
        v = []
        v = [[False for i in range(cols)] for i in range(rows)]

        s_x = -1
        s_y = -1
        
        rotten = []

        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 2:
                    rotten.append((i, j))

                if grid[i][j] == 1:
                    oranges_fresh += 1
    
        q = []
        for i in rotten:
            v[i[0]][i[1]] = True
            q.append((i[0], i[1], 0))
        
        while len(q) > 0:
            
            x = q[0][0]
            y = q[0][1]
            t = q[0][2]
            
            if t > time:
                time = t
            
            q = q[1:]
            
            for i in range(4):
                adj_x = x + self.row_d[i]
                adj_y = y + self.col_d[i]
                
                if is_valid(adj_x, adj_y):
                    oranges_fresh -= 1
                        
                    v[adj_x][adj_y] = True
                    q.append((adj_x, adj_y, t+1))
        
        # print(oranges_fresh) 
        return time if oranges_fresh == 0 else -1

if __name__ == "__main__":
    test_cases = [
        [
            [
                [2,1,1],
                [1,1,0],
                [0,1,1],
            ],
            4
        ],
        [
            [
                [2,1,1],
                [0,1,1],
                [1,0,1],
            ],
            -1,
        ],
        [
            [
                [0,2]
            ],
            0,
        ],
        [
            [[0]],
            0
        ],
        [
            [[1]],
            -1
        ],
        [
            [[0,0,0,0]],
            0
        ],
        [
            [[0,1]],
            -1
        ],
        [
            [[0,2,2]],
            0
        ],
        [
            [[2,1,1],[1,1,1],[0,1,2]],
            2,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().orangesRotting(i[0]))