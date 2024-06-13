# tags: array, sliding_window

from typing import List

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        if len(nums) == 1:
            return nums[0]
        
        c = o = sum(nums[:k])
        for i in range(len(nums) - k):
            c = c - nums[i] + nums[i+k]
            o = max(o, c)
        return o / k
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,12,-5,-6,50,3],
            4,
            12.75,
        ],
        [
            [5],
            1,
            5.00,
        ],
        [
            [0,1,1,3,3],
            4,
            2.00,
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().findMaxAverage(i[0], i[1]))
