# tags: hash_map

from typing import List

class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        h = {}
        for i in arr:
            h[i] = h.get(i, 0) + 1
        return len(h) == len(set(h.values()))
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,2,2,1,1,3],
            True,
        ],
        [
            [1,2],
            False,
        ],
        [
            [-3,0,1,-3,1,1,1,-3,10,0],
            True
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().uniqueOccurrences(i[0]))
        