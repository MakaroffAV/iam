class Token:
    def __init__(self, id, value, code_id, expires, created, is_active):
        self.id = id
        self.value = value
        self.code_id = code_id
        self.expires = expires
        self.created = created
        self.is_active = is_active

    def json(self) -> dict:
        return {
            "id": self.id,
            "value": self.value,
            "code_id": self.code_id,
            "expires": self.expires.isoformat(),
            "created": self.created.isoformat(),
            "is_active": self.is_active,
        }