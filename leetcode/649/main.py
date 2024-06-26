# tags: queue, string

class Solution:
    def __init__(self):
        self.result = {
            "D": "Dire",
            "R": "Radiant",
        }
        self.opposition = {
            "D": "R",
            "R": "D"
        }

    def predictPartyVictory(self, senate: str) -> str:
        s = [*senate]
        
        while True:
            
            if len(set(s)) == 1:
                return self.result[s[0]]
            
            l = s[0]
            for i in range(1, len(s)):
                if s[i] == self.opposition[l]:
                    s.pop(i)
                    s = s[1:]
                    s.append(l)
                    
                    break

if __name__ == "__main__":
    test_cases = [
        [
            "RD",
            "Radiant"
        ],
        [
            "RDD",
            "Dire",
        ],
        [
            "DDRRR",
            "Dire"
        ]
    ]
    for i in test_cases:
        print(i[1] == Solution().predictPartyVictory(i[0]))
