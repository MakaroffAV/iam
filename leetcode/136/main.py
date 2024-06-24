# tags: bit_manipulation

from typing import List

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        a = 0
        for i in nums:
            a ^= i
        return a
    
if __name__ == "__main__":
    test_cases = [
        [
            [2,2,1],
            1,
        ],
        [
            [4,1,2,1,2],
            4,
        ],
        [
            [1],
            1
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().singleNumber(i[0]))
