# tags: design, queue

class RecentCounter:

    def __init__(self) -> None:
        self.q = []

    def ping(self, t: int) -> int:
        while self.q and self.q[0] < t - 3000:
            self.q = self.q[1:]
        self.q.append(t)
        return len(self.q)
    
    
if __name__ == "__main__":
    test_cases = [
        [
            [
                [
                    "init",
                    None,
                ],
                [
                    "ping",
                    1,
                ],
                [
                    "ping",
                    100,
                ],
                [
                    "ping",
                    3001,
                ],
                [
                    "ping",
                    3002,
                ]
            ],
            [
                None,
                1,
                2,
                3,
                3
            ]
        ]
    ]
    
    for case in test_cases:
        r = []
        c = RecentCounter()
        for i in case[0]:
            if i[0] == "init":
                r.append(None)
            else:
                r.append(c.ping(i[1]))
                
        print(r == case[1])
        