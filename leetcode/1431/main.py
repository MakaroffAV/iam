# tags: array, greedy, leetcode

from typing import List

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        m = max(candies)
        r = [False] * len(candies)
        
        for i, v in enumerate(candies):
            if v + extraCandies >= m:
                r[i] = True
        return r
    
if __name__ == "__main__":
    
    test_cases = [
        [
            [2, 3, 5, 1, 3],
            3,
            [True,True,True,False,True],
        ],
        [
            [4, 2, 1, 1, 2],
            1,
            [True,False,False,False,False] 
        ],
        [
            [12, 1, 12],
            10,
            [True,False,True]
        ]
    ]
    
    for i in test_cases:
        print(i[2] == Solution().kidsWithCandies(i[0], i[1]))
