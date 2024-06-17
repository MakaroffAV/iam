# tags: array, binary_search
import math

from typing import List

class Solution:
    def num_of_hours(self, k, piles):
        t = 0
        for p in piles:
            t += math.ceil(p / k)
        return t

    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        
        lo = 1
        k = -1
        hi = max(piles)
        
        while lo <= hi:
            m = lo + ((hi - lo) >> 1)
            if self.num_of_hours(m, piles) <= h:
                k = m
                hi = m - 1
            else:
                lo = m + 1
        return k

if __name__ == "__main__":
    test_cases = [
        [
            [3,6,7,11],
            8,
            4,
        ],
        [
            [30,11,23,4,20],
            5,
            30,
        ],
        [
            [30,11,23,4,20],
            6,
            23,
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().minEatingSpeed(i[0], i[1]))
