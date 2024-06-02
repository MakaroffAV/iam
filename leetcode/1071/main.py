from math import gcd


class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        if str1+str2 != str2+str1:
            return ""
        else:
            return str1[:gcd(len(str1), len(str2))]

        
if __name__ == "__main__":

    test_cases = [
        [
            "ABCABC",
            "ABC",
            "ABC"
        ],
        [
            "ABABAB",
            "ABAB",
            "AB",
        ],
        [
            "LEET",
            "CODE",
            ""
        ]
    ]

    s = Solution()
    for i in test_cases:
        print(i[2] == s.gcdOfStrings(i[0], i[1]))
