import json

from src.service.project import Project as ServiceProject


class Project:
    def __init__(self, service_project: ServiceProject) -> None:
        self.__service_project = service_project

    def all(self, token: str | None, status: str | None):
        if token is None:
            return self.__response(
                {"reason": "Token form field is empty"}, 400, "application/json",
            )
        if status not in ("all", "created", "processing", "ready", "deleted"):
            return self.__response(
                {"reason": "Status value is not allowed"}, 400, "application/json",
            )

        try:
            projects = []
            for i in self.__service_project.all(token, status):
                projects.append(i.json())
            return self.__response({"data": projects}, 200, "application/json")
        except Exception as e:
            print(e)
            return self.__response({"reason": "Unhandled server error occurred"}, 500, "application/json",)

    def get(self, token: str | None, project_uuid: str | None):
        if token is None:
            return self.__response(
                {"reason": "Token form field is empty"}, 400, "application/json",
            )
        if project_uuid is None:
            return self.__response(
                {"reason": "Project uuid form field is empty"}, 400, "application/json",
            )

        try:
            return self.__response(
                self.__service_project.get(token, project_uuid)[0].json(), 200, "application/json",
            )
        except Exception as e:
            print(e)
            return self.__response({"reason": "Unhandled server error occurred"}, 500, "application/json",)

    def create(self, name: str | None, desc: str | None, token: str | None, docs: list):
        if name is None:
            return self.__response(
                {"reason": "Name form field is empty"},  400, "application/json",
            )

        if desc is None:
            return self.__response(
                {"reason": "Desc form field is empty"},  400, "application/json",
            )

        if token is None:
            return self.__response(
                {"reason": "Token form field is empty"}, 400, "application/json",
            )
        if len(docs) > 5:
            return self.__response(
                {"reason": "Doc number limit exceeded"}, 400, "application/json",
            )

        try:
            return self.__response(
                self.__service_project.new(name, desc, token, docs)[0].json(), 200, "application/json",
            )
        except Exception as e:
            print(e)
            return self.__response({"reason": "Unhandled server error occurred"}, 500, "application/json",)

    def new_query(self, token: str | None, query: str | None, enrichment: str | None, project_uuid: str | None):
        if token is None or query is None or enrichment is None or project_uuid is None:
            return self.__response(
                {"reason": "Поля формы: token, query, enrichment; должны быть заполнены"}, 400, "application/json"
            )
        try:
            return self.__response(
                self.__service_project.new_query(token, query, enrichment, project_uuid)[0].json(), 200, "application/json",
            )
        except Exception as e:
            print(e)
            return self.__response({"reason": "Unhandled server error occurred"}, 500, "application/json")

    def query_history(self, token: str | None, project_uuid: str | None):
        if token is None or project_uuid is None:
            return self.__response(
                {"reason": "Поля формы: token, query, enrichment; должны быть заполнены"}, 400, "application/json"
            )
        try:
            query_history = []
            for i in self.__service_project.query_history(token, project_uuid):
                query_history.append(i.json())
            return self.__response({"data": query_history}, 200, "application/json")
        except Exception as e:
            print(e)
            return self.__response({"reason": "Unhandled server error occurred"}, 500, "application/json")

    def __response(self, body: dict, status_code: int, content_type: str):
        return {
            "response": json.dumps(body), "status": status_code, "mimetype": content_type
        }

