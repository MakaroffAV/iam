# tags: array, two_pointers, binary_search, sorting

import math
from typing import List

class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        res = []
        potions.sort()
        
        for i in spells:
            l, r = 0, len(potions)-1
            while l <= r:
                mid = (l + r) // 2
                if potions[mid] * i >= success:
                    r = mid - 1
                else:
                    l = mid + 1
            res.append(len(potions) - l)
        return res

if __name__ == "__main__":
    test_cases = [
        [
            [5,1,3],
            [1,2,3,4,5],
            7,
            [4,0,3],
        ],
        [
            [3,1,2],
            [8,5,8],
            16,
            [2,0,2],
        ]
    ]
    for i in test_cases:
        print(i[3] == Solution().successfulPairs(i[0], i[1], i[2]))
