# tags: array, matrix, hash_table

from typing import List

class Solution:
    def equalPairs(self, grid: List[List[int]]) -> int:
        c, n = 0,  len(grid)
        t = list(zip(*grid))
        for i in range(n):
            for j in range(n):
                if grid[i] == list(t[j]):
                    c += 1
        return c
    
if __name__ == "__main__":
    
    test_cases = [
        [
            [[3,2,1],[1,7,6],[2,7,7]],
            1
        ],
        [
            [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]],
            3
        ],
        [
            [[13,13],[13,13]],
            4
        ]
    ]
    
    for i in test_cases:
        print(i[1] == Solution().equalPairs(i[0]))
        