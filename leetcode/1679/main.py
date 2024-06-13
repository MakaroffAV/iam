# tags: array, hash_table, two_pointers

from typing import List

class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        
        nums.sort()
        
        a = 0
        l = 0
        r = len(nums) - 1
        while l < r:
            temp = nums[l] + nums[r]
            if temp  == k:
                a += 1
                l += 1
                r -= 1
            elif temp > k:
                r -= 1
            else:
                l += 1
        
        return a
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,3,4],
            5,
            2
        ],
        [
            [3,1,3,4,3],
            6,
            1
        ],
        [
            [4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4],
            2,
            2
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().maxOperations(i[0], i[1]))
