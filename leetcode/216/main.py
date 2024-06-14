# tags: array, backtracking

from typing import List

class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        r = []
        
        def dfs(start, path, curr_sum):
            if   len(path) == k and curr_sum == n:
                r.append(path)
                return
            elif len(path) == k and curr_sum != n:
                return
            
            for i in range(start, 10):
                dfs(i+1, path+[i], curr_sum+i)
        
        dfs(1, [], 0)
        return r
    
if __name__ == "__main__":
    test_cases = [
        [
            3,
            7,
            [[1,2,4]],
        ],
        [
            3,
            9,
            [[1,2,6],[1,3,5],[2,3,4]]
        ],
        [
            4,
            1,
            [],
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().combinationSum3(i[0], i[1]))
