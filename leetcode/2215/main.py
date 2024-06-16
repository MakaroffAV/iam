# tags: hash_table, array

from typing import List

class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        h1 = {
            i: True
            for i in nums1
        }
        h2 = {
            i: True
            for i in nums2
        }
        
        r1 = []
        r2 = []
        
        for i in h1:
            if not h2.get(i, False):
                r1.append(i)
                
        for i in h2:
            if not h1.get(i, False):
                r2.append(i)
                
        return [r1, r2]

if __name__ == "__main__":
    test_cases = [
        [
            [1, 2, 3],
            [2, 4, 6],
            [
                [1, 3],
                [4, 6]
            ]
        ],
        [
            [1, 2, 3, 3],
            [1, 1, 2, 2],
            [
                [3],
                [],
            ]
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().findDifference(i[0], i[1]))
