import uuid as uuid_lib

from src.repo.code import Code as RepoCode
from src.repo.token import Token as RepoToken

from src.domain.token import Token as DomainToken


class Token:
    def __init__(self, repo_code: RepoCode, repo_token: RepoToken) -> None:
        self.__repo_code = repo_code
        self.__repo_token = repo_token

    def check(self, token: str) -> bool:
        is_active \
            = self.__repo_token.is_active(token)
        return False if is_active is None else is_active

    def create(self, code: str, uuid: str) -> DomainToken:
        new_uuid, new_code \
            = uuid_lib.uuid4(), self.__repo_code.get_by_code_and_uuid(code, uuid)

        return self.__repo_token.create(new_code[0].id, str(new_uuid))
