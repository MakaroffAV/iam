from src.config.db import Database
from src.domain.code import Code as DomainCode

SQL_CODE_CREATE = """
insert into ai_app.code(user_id, code, hash) values (%s, %s, %s)
"""

SQL_CODE_GET_BY_UUID = """
select id, hash, code, user_id, created from ai_app.code where hash = %s
"""

SQL_CODE_GET_BY_CODE_AND_UUID = """
select id, hash, code, user_id, created from ai_app.code where hash = %s and code = %s
"""


class Code:

    def __init__(self, db: Database):
        self.__db = db

    def get_by_uuid(self, uuid: str):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_CODE_GET_BY_UUID, (uuid,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainCode(i[0], i[1], i[2], i[3], i[4]))
        return res if len(res) != 0 else None

    def create(self, user_id: int, code: int, uuid: str):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_CODE_CREATE, (user_id, code, uuid,)
                )
        return self.get_by_uuid(uuid)

    def get_by_code_and_uuid(self, code: str, uuid: str):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_CODE_GET_BY_CODE_AND_UUID, (uuid, code,),
                )
                for i in db_curs.fetchall():
                    res.append(DomainCode(i[0], i[1], i[2], i[3], i[4]))
        return res if len(res) != 0 else None