from datetime import datetime


class ProjectQuery:
    def __init__(
            self,
            id: int,
            uuid: str,
            project_id: int,
            query: str,
            enrichment: bool,
            status: str,
            created) -> bool:

        self.id = id
        self.uuid = uuid
        self.project_id = project_id
        self.query = query
        self.enrichment = enrichment
        self.status = status
        self.created = created

    def json(self):
        return {
            "id": self.id,
            "uuid": self.uuid,
            "project_id": self.project_id,
            "query": self.query,
            "enrichment": self.enrichment,
            "status": self.status,
            "created": self.created.strftime("%Y-%m-%d %H:%M:%S"),
        }
