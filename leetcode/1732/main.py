# tags: prefix_sum, array

from typing import List

class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        m = 0
        c = 0
        for i in gain:
            c += i
            if c > m:
                m = c
        return m
    
if __name__ == "__main__":
    test_cases = [
        [
            [-5,1,5,0,-7],
            1,
        ],
        [
            [-4,-3,-2,-1,4,3,2],
            0
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().largestAltitude(i[0]))
