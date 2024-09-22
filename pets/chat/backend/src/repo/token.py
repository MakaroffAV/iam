from src.domain.user import User as DomainUser
from src.domain.token import Token as DomainToken

SQL_TOKEN_CREATE = """
insert into ai_app.token(value, code_id) values (%s, %s)
"""

SQL_TOKEN_IS_ACTIVE = """
select exists (select id from ai_app.token where value = %s and is_active = true and now() < expires)
"""

SQL_TOKEN_GET_BY_UUID = """
select id, value, code_id, expires, created, is_active from ai_app.token where value = %s
"""

SQL_TOKEN_GET_USER_BY_TOKEN = """
select
    t3.id, t3.email, t3.created
from
    ai_app.token as t1
left join
    ai_app.code  as t2
on
    t1.code_id = t2.id
left join
    ai_app.user  as t3
on
    t2.user_id = t3.id
where
    t1.value = %s and t1.is_active = true and now() < t1.expires
"""


class Token:
    def __init__(self, db):
        self.__db = db

    def is_active(self, token: str) -> bool | None:
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_TOKEN_IS_ACTIVE, (token,),
                )
                return db_curs.fetchone()[0]

        return None

    def get_by_uuid(self, uuid: str) -> list[DomainToken] | None:
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_TOKEN_GET_BY_UUID, (uuid,)
                )
                for i in db_curs.fetchall():
                    print(i)
                    res.append(DomainToken(i[0], i[1], i[2], i[3], i[4], i[5]))
        return res if len(res) != 0 else None

    def create(self, code_id: int, uuid: str) -> list[DomainToken] | None:
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_TOKEN_CREATE, (uuid, code_id,),
                )
        return self.get_by_uuid(uuid)

    def get_user_by_token(self, token: str):
        res = []
        print("repo", token)
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_TOKEN_GET_USER_BY_TOKEN, (token,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainUser(*i))
        return res if len(res) != 0 else None
