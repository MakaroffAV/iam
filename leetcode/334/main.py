# tags: array, greedy

from typing import List

class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        n = len(nums)
        
        if n < 3:
            return False
        
        f = float("inf")
        s = float("inf")
        for i in nums:
            if i <= f:
                f = i
            elif i <= s:
                s = i
            else:
                return True

        return False
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,3,4,5],
            True,
        ],
        [
            [5,4,3,2,1],
            False,
        ],
        [
            [2,1,5,0,4,6],
            True,
        ],
        [
            [1,5,0,4,1,3],
            True,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().increasingTriplet(i[0]))
