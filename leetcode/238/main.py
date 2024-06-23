# tags: prefix_sum, array

from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        l = 0
        r = n-1
        
        l_prod = [0]*n
        r_prod = [0]*n
        
        while r >= 0:
            l_prod[l] = nums[l] if l == 0   else nums[l] * l_prod[l-1]
            r_prod[r] = nums[r] if r == n-1 else nums[r] * r_prod[r+1]
            l += 1
            r -= 1
            
        res     = [0] * n
        res[0]  = r_prod[1]
        res[-1] = l_prod[n-2]
        for i in range(1, n-1):
            res[i] = l_prod[i-1] * r_prod[i+1]
        return res
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,3,4],
            [24,12,8,6],
        ],
        [
            [-1,1,0,-3,3],
            [0,0,9,0,0],
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().productExceptSelf(i[0]))
