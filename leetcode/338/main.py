# tags: dp, bit_manipulation

from typing import List

class Solution:
    def countBits(self, n: int) -> List[int]:
        r = []
        for i in range(n+1):
            r.append(i.bit_count())
        return r
    
if __name__ == "__main__":
    test_cases = [
        [
            2,
            [0,1,1]
        ],
        [
            5,
            [0,1,1,2,1,2]
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().countBits(i[0]))
    