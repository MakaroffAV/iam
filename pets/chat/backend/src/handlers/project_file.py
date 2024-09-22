import json

from src.service.project_file import ProjectFile as ServiceProjectFile


class ProjectFile:
    def __init__(
            self,
            service_project_file: ServiceProjectFile) -> None:

        self.__service_project_file = service_project_file

    def list(
            self,
            token: str | None,
            project_uuid: str | None) -> dict:

        if token is None or project_uuid is None:
            return self.__response(
                400, "application/json", {"reason": "Некорректная форма запроса"}
            )

        try:
            project_files = []
            for i in self.__service_project_file.list(token, project_uuid):
                project_files.append(i.json())
            return self.__response(200, "application/json", {"data": project_files})
        except Exception as e:
            print(e)
            return self.__response(
                500, "application/json", {"reason": "Произошла внутренняя ошибка сервера"}
            )

    def __response(self, status_code: int, content_type: str, body: dict) -> dict:
        return {
            "response": json.dumps(body), "status": status_code, "mimetype": content_type
        }
