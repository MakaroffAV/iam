class ProjectFile:
    def __init__(self, id, project_id, uuid, name, body, created):
        self.id = id
        self.uuid = uuid
        self.name = name
        self.body = body
        self.created = created
        self.project_id = project_id

    def json(self) -> dict:
        return {
            "id": self.id,
            "uuid": self.uuid,
            "name": self.name,
            "body": self.body,
            "created": self.created.isoformat(),
            "project_id": self.project_id,
        }
