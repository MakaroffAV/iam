from src.config.db import Database
from src.domain.user import User as DomainUser

SQL_USER_CREATE = """
insert into ai_app.user (email) values (%s)
"""

SQL_USER_EXISTS = """
select exists (select id from ai_app.user where email = %s)
"""

SQL_USER_GET_BY_EMAIL = """
select id, email, created from ai_app.user where email = %s
"""

SQL_USER_GET_BY_ID = """
select id, email, created from ai_app.user where id = %s
"""


class User:

    def __init__(self, db: Database):
        self.__db = db

    def exists(self, email):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_USER_EXISTS, (email,)
                )
                return  db_curs.fetchone()[0]

    def create(self, email: str):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_USER_CREATE, (email,)
                )
        return self.get_by_email(email)

    def get_by_id(self, id: int):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_USER_GET_BY_ID, (id,),
                )
                for i in db_curs.fetchall():
                    res.append(DomainUser(*i))
        return res

    def get_by_email(self, email: str):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_USER_GET_BY_EMAIL, (email,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainUser(i[0], i[1], i[2]))
        return res if len(res) != 0 else None
