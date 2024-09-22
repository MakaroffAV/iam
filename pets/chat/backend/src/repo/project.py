from src.domain.project import Project as DomainProject


SQL_PROJECT_GET_BY_UUID = """
select id, uuid, name, status, user_id, created, description from ai_app.project where uuid = %s
"""

SQL_PROJECT_CREATE = """
insert into ai_app.project(name, description, uuid, user_id) values (%s, %s, %s, %s)
"""

SQL_PROJECT_GET_BY_ID = """
select id, uuid, name, status, user_id, created, description from ai_app.project where id = %s 
"""

SQL_PROJECT_GET_BY_USER_ID = """
select id, uuid, name, status, user_id, created, description from ai_app.project where user_id = %s order by created desc 
"""

SQL_PROJECT_UPDATE_STATUS_BY_ID = """
update ai_app.project set status = %s where id = %s
"""

SQL_PROJECT_GET_BY_USER_ID_WITH_STATUS = """
select id, uuid, name, status, user_id, created, description from ai_app.project where user_id = %s and status = %s order by created desc
"""


class Project:
    def __init__(self, db):
        self.__db = db

    def get_by_id(self, id: int):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_GET_BY_ID, (id,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProject(*i))
        return res

    def get_by_uuid(self, uuid: str):
        res = []
        with self.__db as db_conn:
            print(uuid)
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_GET_BY_UUID, (uuid,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProject(*i))
        return res if len(res) != 0 else None

    def create(self, name: str, desc: str, uuid: str, user_id: int):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_CREATE, (name, desc, uuid, user_id,),
                )
        return self.get_by_uuid(uuid)

    def get_by_user_id(self, user_id: int, status: str):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                if status == "all":
                    db_curs.execute(
                        SQL_PROJECT_GET_BY_USER_ID, (user_id,),
                    )
                else:
                    db_curs.execute(
                        SQL_PROJECT_GET_BY_USER_ID_WITH_STATUS, (user_id, status.title(),)
                    )
                for i in db_curs.fetchall():
                    res.append(DomainProject(*i))
        return res

    def update_status_by_id(self, project_id: int, status: str):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_UPDATE_STATUS_BY_ID, (status, project_id,),
                )
        return self.get_by_id(project_id)
