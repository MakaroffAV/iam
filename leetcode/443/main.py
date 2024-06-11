# tags: strings, two_pointers

from typing import List

class Solution:
    def update(self, l, r, c) -> str:
        if r - l == 0:
            return "{}".format(c)
        else:
            return "{}{}".format(c, r-l+1)

    def compress(self, chars: List[str]) -> int:
        if len(chars) == 1:
            return 1

        l = 0
        r = 0
        a = ""
        while r < len(chars) - 1:
            if chars[r+1] == chars[r]:
                r += 1
            else:
                a += self.update(l, r, chars[l])
                r += 1
                l  = r

        a += self.update(l, r, chars[l])
        
        chars[:len(a)]  = [*a]
        chars = chars[:len(a)]         
        return len(chars)
    
if __name__ == "__main__":
    test_cases = [
        [
            ["a","a","b","b","c","c","c"],
            6,
        ],
        [
            ["a"],
            1,
        ],
        [
            ["a","b","b","b","b","b","b","b","b","b","b","b","b"],
            4,
        ],
        [
            ["a","a","a","b","b","a","a"],
            6,
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().compress(i[0]))
