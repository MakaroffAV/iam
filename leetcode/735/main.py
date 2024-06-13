# tags: array, stack

from typing import List

class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        stack = []
        for i in asteroids:
            while stack and stack[-1] > 0 and i < 0:
                if stack[-1]  == -i:
                    stack.pop()
                    break
                elif stack[-1] < -i:
                    stack.pop()
                    continue
                else:
                    break
            else:
                stack.append(i)

        return stack

if __name__ == "__main__":
    test_cases = [
        [
            [5,10,-5],
            [5,10], 
        ],
        [
            [8,-8],
            []
        ],
        [
            [10,2,-5],
            [10]
        ],
        [
            [-2,-1,1,2],
            [-2,-1,1,2]
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().asteroidCollision(i[0]))
