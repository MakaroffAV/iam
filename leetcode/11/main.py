# tags: array, greedy, two_pointers

from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        l = 0
        s = 0
        r = len(height) - 1
        
        while l < r:
            s = max(min(height[l], height[r]) * (r - l), s)
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        return s
    
if __name__ == "__main__":
    
    test_cases = [
        [
            [1,8,6,2,5,4,8,3,7],
            49,
        ],
        [
            [1,1],
            1
        ]
    ]
    
    for i in test_cases:
        print(i[1] == Solution().maxArea(i[0]))
