# tags: string, two_pointers

class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(s.split()[::-1])

if __name__ == "__main__":
    test_cases = [
        [
            "the sky is blue",
            "blue is sky the",
        ],
        [
            "  hello world  ",
            "world hello",
        ],
        [
            "a good   example",
            "example good a",
        ]
    ]
    
    for i in test_cases:
        print(i[1] == Solution().reverseWords(i[0]))
