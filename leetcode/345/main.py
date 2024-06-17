# tags: string, two_pointers

class Solution:
    def reverseVowels(self, s: str) -> str:
        s = [*s]
        l = 0
        r = len(s) - 1

        v = {
            "a": True,
            "e": True,
            "i": True,
            "o": True,
            "u": True,
        }
        
        while l < r:
            c_l = s[l].lower()
            c_r = s[r].lower()

            if v.get(c_l, False) and v.get(c_r, False):
                s[l], s[r] = s[r], s[l]
                l += 1
                r -= 1
                continue
            
            if v.get(c_l, False) and not v.get(c_r, False):
                r -= 1
                continue
            
            if not v.get(c_l, False) and v.get(c_r, False):
                l += 1
                continue
            
            if not v.get(c_l, False) and not v.get(c_r, False):
                l += 1
                r -= 1
                continue

        return "".join(s)

if __name__ == "__main__":
    test_cases = [
        [
            "hello",
            "holle"
        ],
        [
            "leetcode",
            "leotcede"
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().reverseVowels(i[0]))
