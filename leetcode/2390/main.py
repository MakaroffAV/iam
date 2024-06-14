# tags: stack, string

class Solution:
    def removeStars(self, s: str) -> str:
        stack = []
        for i in s:
            if i == "*":
                stack.pop()
            else:
                stack.append(i)
        return "".join(stack)
    
if __name__ == "__main__":
    test_cases = [
        [
            "leet**cod*e",
            "lecoe",
        ],
        [
            "erase*****",
            ""
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().removeStars(i[0]))
