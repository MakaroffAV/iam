# tags: array, prefix_sum

from typing import List

class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        n = len(nums)
       
        l = 0
        r = n - 1
         
        l_sum = [0] * n
        r_sum = [0] * n
        
        while r >= 0:
            l_sum[l] = nums[l] if l == 0   else nums[l] + l_sum[l-1]
            r_sum[r] = nums[r] if r == n-1 else nums[r] + r_sum[r+1]
            l += 1
            r -= 1
            
        for i in range(n):
            if l_sum[i] == r_sum[i]:
                return i
        return -1
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,7,3,6,5,6],
            3,
        ],
        [
            [1,2,3],
            -1,
        ],
        [
            [2,1,-1],
            0,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().pivotIndex(i[0]))
        