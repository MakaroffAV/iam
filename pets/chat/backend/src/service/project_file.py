import uuid

from src.repo.project_file import ProjectFile as RepoProjectFile
from src.domain.project_file import ProjectFile as DomainProjectFile


class ProjectFile:
    def __init__(self, repo_project_file: RepoProjectFile):
        self.__repo_project_file = repo_project_file

    def create(self, name, body, project_id):
        new_uuid = str(uuid.uuid4())
        return self.__repo_project_file.create(name, body, project_id, new_uuid)

    def list(self, token: str, project_uuid: str):
        return self.__repo_project_file.get_meta_by_project_uuid(project_uuid)

    def get_by_id(self, id: int) -> DomainProjectFile | None:
        return self.__repo_project_file.get_by_id(id)

    def get_meta_by_project_id(self, project_id):
        return self.__repo_project_file.get_meta_by_project_id(project_id)