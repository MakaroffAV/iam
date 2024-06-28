# tags: dp

from typing import List

class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n = len(cost)
        for i in range(n-3, -1, -1):
            cost[i] += min(cost[i+1], cost[i+2])
        return min(cost[0], cost[1])
    
if __name__ == "__main__":
    test_cases = [
        [
            [10,15,20],
            15,
        ],
        [
            [1,100,1,1,1,100,1,1,100,1],
            6
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().minCostClimbingStairs(i[0]))
