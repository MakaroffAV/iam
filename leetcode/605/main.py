# tags: array, greedy

from typing import List

class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        if n == 0:
            return True
        for i in range(len(flowerbed)):
            if flowerbed[i] == 0 and (i == 0 or flowerbed[i-1] == 0) and (i == len(flowerbed)-1 or flowerbed[i+1] == 0):
                flowerbed[i] = 1
                n -= 1
                if n == 0:
                    return True
        return False
    
if __name__ == "__main__":
    test_cases = [
        [
            [1,0,0,0,1],
            1,
            True,
        ],
        [
            [1,0,0,0,1],
            2,
            False
        ],
        [
            [1,0,0,0,0,0,1],
            2,
            True,
        ],
        [
            [1,0,1,0,1,0,1],
            0,
            True
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().canPlaceFlowers(i[0], i[1]))
        