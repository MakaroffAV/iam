# tags: stack, string, recursion

class Solution:

    def decodeString(self, s: str) -> str:
        stack = []
        for char in s:
            if char == "]":
                r = ""
                while stack[-1] != "[":
                    r += stack.pop()
                stack.pop()
                n = ""
                while len(stack) != 0 and stack[-1].isdigit() == True:
                    n += stack.pop()
                stack.append(r * int(n[::-1]))
            else:
                stack.append(char)
                
        return "".join([i[::-1] for i in stack])
    
if __name__ == "__main__":
    test_cases = [
        [
            "3[a]2[bc]",
            "aaabcbc",
        ],
        [
            "3[a2[c]]",
            "accaccacc",
        ],
        [
            "2[abc]3[cd]ef",
            "abcabccdcdcdef",
        ]
    ]
    for case in test_cases:
        print(case[1] == Solution().decodeString(case[0]))
