# tags: array, binary_search

from typing import List

class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        
        if len(nums) == 1:
            return 0
        
        if nums[0] > nums[1]:
            return 0
        
        if nums[-1] > nums[-2]:
            return len(nums) - 1

        l = 0
        r = len(nums) - 1
        
        while l <= r:
            m = (l + r) // 2
            if nums[m] > nums[m+1] and nums[m] > nums[m-1]:
                return m
            if nums[m] <= nums[m+1]:
                l = m + 1
            else:
                r = m - 1
        
        return -1
    
if __name__ == "__main__":
    
    test_cases = [
        [
            [1,2,3,1],
            2,
        ],
        [
            [1,2,1,3,5,6,4],
            5,
        ],
        [
            [3,4,3,2,1],
            1
        ]
    ]
    for test_case in test_cases:
        print(test_case[1] == Solution().findPeakElement(test_case[0]))
        