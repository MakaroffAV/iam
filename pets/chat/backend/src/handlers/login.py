import json

from src.service.login import Login as ServiceLogin


class Login:

    def __init__(self, sv_login: ServiceLogin) -> None:
        self.__sv_login = sv_login

    def post(self, email: str | None):
        if email is None:
            return self.__response(
                {"reason": "Email form field is empty"}, 400, "application/json",
            )

        try:
            code = self.__sv_login.login(email)
            return self.__response(code[0].json(), 200, "application/json")
        except Exception as e:
            print(e)
            return self.__response(
                {"reason": "Unhandled server error occurred"}, 500, "application/json",
            )

    def __response(self, body: dict, status_code: int, content_type: str):
        return {
            "response": json.dumps(body), "status": status_code, "mimetype": content_type
        }
