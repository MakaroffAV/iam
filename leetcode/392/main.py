# tags: string, two_pointers

class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        sp = tp = 0
        while sp < len(s) and tp < len(t):
            if s[sp] == t[tp]:
                sp += 1
            tp += 1
        return sp == len(s)
    
if __name__ == "__main__":
    test_cases = [
        [
            "abc",
            "ahbgdc",
            True,
        ],
        [
            "axc",
            "ahbgdc",
            False,
        ]
    ]
    for i in test_cases:
        print(i[2] == Solution().isSubsequence(i[0], i[1]))
