# tags: hash_table, string, trie

class Trie:

    def __init__(self):
        self.root={}
        
    def insert(self, word: str) -> None:
        cur = self.root
        for ch in word:
            if ch not in cur:
                cur[ch] = {}
            cur = cur[ch]
        cur["*"] = ""

    def search(self, word: str) -> bool:
        cur = self.root
        for ch in word:
            if ch not in cur:
                return False
            cur = cur[ch]
        return "*" in cur
        
    def startsWith(self, prefix: str) -> bool:
        cur  = self.root
        for l in prefix:
            if l not in cur:
                return False
            cur = cur[l]
        return True
        
if __name__ == "__main__":
    t = Trie()
    c = [
        [
            t.insert,
            "apple"
        ],
        [
            t.search,
            "apple",
        ],
        [
            t.search,
            "app"
        ],
        [
            t.startsWith,
            "app"
        ],
        [
            t.insert,
            "app",
        ],
        [
            t.search,
            "app"
        ]
    ]
    for i in c:
        print(i[0](i[1]))
