import json
from src.service.token import Token as ServiceToken


class Token:

    def __init__(self, sv_token: ServiceToken) -> None:
        self.__sv_token = sv_token

    def check(self, token: str | None):
        if token is None:
            return self.__response(
                {"reason": "token form field is empty"}, 400, "application/json",
            )

        try:
            return self.__response(
                {"ok": self.__sv_token.check(token)}, 200, "application/json",
            )
        except Exception as e:
            print(e)
            return self.__response(
                {"reason": "Unhandled server error occurred"}, 500, "application/json",
            )

    def create(self, code: str | None, uuid: str | None):
        if code is None or uuid is None:
            return self.__response(
                {"reason": "code or uuid form fields are empty"}, 400, "application/json",
            )
        try:
            token = self.__sv_token.create(code, uuid)
            return self.__response(token[0].json(), 200, "application/json")
        except Exception as e:
            print(e)
            return self.__response(
                {"reason": "Unhandled server error occurred"}, 500, "application/json",
            )

    def __response(self, body: dict, status_code: int, content_type: str):
        return {
            "response": json.dumps(body), "status": status_code, "mimetype": content_type
        }