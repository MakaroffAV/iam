class Code:
    def __init__(self, id, hash, code, user_id, created):
        self.id = id
        self.hash = hash
        self.code = code
        self.user_id = user_id
        self.created = created

    def json(self) -> dict:
        return {
            "id": self.id,
            "hash": self.hash,
            "code": self.code,
            "user_id": self.user_id,
            "created": self.created.isoformat(),
        }
