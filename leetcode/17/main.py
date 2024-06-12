# tags: hash_table, string, backtracking

from typing import List

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        
        keys = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        
        def backtrack(c, n):
            if len(n) == 0:
                r.append(c)
            else:
                for i in keys[n[0]]:
                    backtrack(c + i, n[1:])
        
        if len(digits) == 0:
            return []
        if len(digits) == 1:
            return [*keys[digits]]
        
        r = []
        backtrack("", digits)

        return r
    
if __name__ == "__main__":
    test_cases = [
        [
            "23",
            ["ad","ae","af","bd","be","bf","cd","ce","cf"],
        ],
        [
            "",
            [],
        ],
        [
            "2",
            ["a","b","c"]
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().letterCombinations(i[0]))
