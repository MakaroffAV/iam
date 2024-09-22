import json
import uuid

from modules.queue import Queue as ModuleQueue

from src.repo.token import Token as RepoToken
from src.repo.project import Project as RepoProject
from src.repo.project_query import ProjectQuery as RepoProjectQuery

from src.service.project_file import ProjectFile as ServiceProjectFile


class Project:
    def __init__(
            self,
            md_queue: ModuleQueue,
            rp_token: RepoToken,
            rp_project: RepoProject,
            sv_project_file: ServiceProjectFile,
            rp_project_query: RepoProjectQuery) -> None:

        self.__md_queue = md_queue
        self.__rp_token = rp_token
        self.__rp_project = rp_project
        self.__sv_project_file = sv_project_file
        self.__rp_project_query = rp_project_query

    def all(self, token: str, status: str | None):
        user = self.__rp_token.get_user_by_token(token)
        return self.__rp_project.get_by_user_id(user[0].id, status)

    def get(self, token: str, project_uuid: str):
        user, project \
            = self.__rp_token.get_user_by_token(token), self.__rp_project.get_by_uuid(project_uuid)
        return None if user[0].id != project[0].user_id else project

    def update_status(self, project_id: int, status: str):
        return self.__rp_project.update_status_by_id(project_id, status)[0]

    def new(self, name: str, desc: str, user_token: str, docs):
        project_uuid = str(uuid.uuid4())
        project_user = self.__rp_token.get_user_by_token(user_token)

        new_project \
            = self.__rp_project.create(name, desc, project_uuid, project_user[0].id)
        if new_project is None:
            return None
        else:
            for i in docs:
                self.__sv_project_file.create(
                    i.filename, i.read(), new_project[0].id,
                )
            self.__md_queue.send(
                "ai_app",
                json.dumps(
                    {
                        "task": "create_project",
                        "data": {
                            "project": new_project[0].json(),
                        },
                    },
                ),
            )

        return new_project

    def new_query(self, token: str, query: str, enrichment: str, project_uuid: str):
        project \
            = self.__rp_project.get_by_uuid(project_uuid)
        if project is None:
            return None
        else:
            new_query = self.__rp_project_query.create(
                project[0].id, query, enrichment == "on", str(uuid.uuid4()),
            )
            print("here1", new_query)
            self.__md_queue.send(
                "ai_app",
                json.dumps(
                    {
                        "task": "new_project_query",
                        "data": new_query[0].json(),
                    }
                )
            )
            print("here2", new_query)
        return new_query

    def query_history(self, token: str, project_uuid: str):
        project \
            = self.__rp_project.get_by_uuid(project_uuid)
        if project is None:
            return None
        return self.__rp_project_query.get_by_project_id(project[0].id)

