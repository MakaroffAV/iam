# tags: dp, array

from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        
        if n == 0:
            return 0
        
        if n == 1:
            return nums[0]
        
        dp = [0] * n
        
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])
        
        for i in range(2, n):
            dp[i] = max(dp[i-1], dp[i-2] + nums[i])
            
        return max(dp[n-1], dp[n-2])

    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,3,1],
            4,
        ],
        [
            [2,7,9,3,1],
            12,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().rob(i[0]))
    