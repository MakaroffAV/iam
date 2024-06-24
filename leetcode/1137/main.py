# tags: dp

class Solution:
    def tribonacci(self, n: int) -> int:
        t0 = 0
        t1 = 1
        t2 = 1
        
        if n == 0:
            return t0
        if n == 1:
            return t1
        if n == 2:
            return t2
        
        for _ in range(3, n+1):
            t0, t1, t2 = t1, t2, t0+t1+t2
        return t2
    
if __name__ == "__main__":
    test_cases = [
        [
            4,
            4,
        ],
        [
            25,
            1389537
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().tribonacci(i[0]))
