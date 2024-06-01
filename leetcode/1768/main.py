class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        
        i = 0
        j = 0
        r = []
        n = len(word1)
        m = len(word2)
        
        while i < n or j < m:
            if i < n:
                r.append(word1[i])
                i += 1
            if j < m:
                r.append(word2[j])
                j += 1

        return "".join(r)
        
        
if __name__ == "__main__":
    
    test_cases = [
        [
            "abc",
            "pqr",
            "apbqcr"
        ],
        [
            "ab",
            "pqrs",
            "apbqrs",
        ],
        [
            "abcd",
            "pq",
            "apbqcd"
        ]
    ]
    
    s = Solution()
    for i in test_cases:
        print(i[2] == s.mergeAlternately(i[0], i[1]))
    