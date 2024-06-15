# tags: array, binary_search, sliding_window

from typing import List

class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        l = r = 0
        for r in range(len(nums)):
            if nums[r] == 0:
                k -= 1
            if k < 0:
                if nums[l] == 0:
                    k += 1
                l += 1
        return r - l + 1
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,1,1,0,0,0,1,1,1,1,0],
            2,
            6,
        ],
        [
            [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1],
            3,
            10
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().longestOnes(i[0], i[1]))
