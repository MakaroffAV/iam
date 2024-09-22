from datetime import datetime


class Project:
    def __init__(
            self,
            id: int, uuid: str, name: str,
            status: str, user_id: int, created, description: str) -> None:

        self.id = id
        self.uuid = uuid
        self.name = name
        self.status = status
        self.user_id = user_id
        self.created = created
        self.description = description

    def json(self):
        return {
            "id": self.id,
            "uuid": self.uuid,
            "name": self.name,
            "status": self.status,
            "user_id": self.user_id,
            "created": self.created.isoformat(),
            "description": self.description
        }
