# tags: string, sliding_window

class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        r = 0
        m = {
            "a": True,
            "e": True,
            "i": True,
            "o": True,
            "u": True,
        }
        
        if len(s) <= k:
            for i in s:
                if m.get(i, False):
                    r += 1
            return r
        
        for i in s[:k]:
            if m.get(i, False):
                r += 1
        t = r       
        for i in range(1, len(s)-k+1):
            if m.get(s[i-1],   False):
                t -= 1
            if m.get(s[i+k-1], False):
                t += 1
            # print(s[i:i+k], s[i-1], s[i+k-1], t)
            if t > r:
                r = t
        return r

if __name__ == "__main__":
    test_cases = [
        [
            "aa",
            3,
            2,
        ],
        [
            "aeiou",
            3,
            3,
        ],
        [
            "aeiou",
            2,
            2,
        ],
        [
            "leetcode",
            3,
            2,
        ],
        [
            "abciiidef",
            3,
            3
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().maxVowels(i[0], i[1]))
