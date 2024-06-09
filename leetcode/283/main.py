# tags: array, two_pointers

from typing import List

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        
        l = 0
        for r in range(len(nums)):
            if nums[r] != 0:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1

        return nums
     
if __name__ == "__main__":
    
    test_cases = [
        [
            [0,1,0,3,12],
            [1,3,12,0,0]
        ],
        [
            [0],
            [0],
        ]
    ]
    
    for i in test_cases:
        print(i[1] == Solution().moveZeroes(i[0]))
