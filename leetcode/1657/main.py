# tags: hash_table, string, sorting, counting

class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        
        if len(word1) != len(word2):
            return False
        
        freq1 = {}
        freq2 = {}
        for i in range(len(word1)):
            freq1[word1[i]] = freq1.get(word1[i], 0) + 1
            freq2[word2[i]] = freq2.get(word2[i], 0) + 1
        
        return set(word1) == set(word2) and sorted(freq1.values()) == sorted(freq2.values())


if __name__ == "__main__":
    
    test_cases = [
        [
            "abc",
            "bca",
            True,
        ],
        [
            "a",
            "aa",
            False,
        ],
        [
            "cabbba",
            "abbccc",
            True
        ],
        [
            "uau",
            "ssx",
            False
        ]
    ]
    
    for i in test_cases:
        print(i[2] == Solution().closeStrings(i[0], i[1]))
