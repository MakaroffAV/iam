# tags: array, dp, sliding_window

from typing import List

class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        m = 0
        l = 0
        z = 0
        n = len(nums)
        for r in range(n):
            if nums[r] == 0:
                z += 1
            while z > 1:
                if nums[l] == 0:
                    z -= 1
                l += 1
            m = max(m, r - l)
        return m
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,1,0,1],
            3,
        ],
        [
            [0,1,1,1,0,1,1,0,1],
            5,
        ],
        [
            [1,1,1],
            2
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().longestSubarray(i[0]))
        