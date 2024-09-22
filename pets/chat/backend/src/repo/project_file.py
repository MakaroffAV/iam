from src.domain.project_file import ProjectFile as DomainProjectFile


SQL_PROJECT_FILE_CREATE = """
insert into ai_app.project_file (name, body, project_id, uuid) values (%s, %s, %s, %s)
"""

SQL_PROJECT_FILE_GET_BY_ID = """
select id, project_id, uuid, name, body, created from ai_app.project_file where id = %s
"""

SQL_PROJECT_FILE_GET_BY_UUID = """
select id, project_id, uuid, name, body, created from ai_app.project_file where uuid = %s
"""

SQL_PROJECT_FILE_GET_META_BY_PROJECT_ID = """
select id, project_id, uuid, name, '', created from ai_app.project_file where project_id = %s
"""

SQL_PROJECT_FILE_GET_META_BY_PROJECT_UUID = """
select t1.id, t1.project_id, t1.uuid, t1.name, '', t1.created from ai_app.project_file as t1 left join ai_app.project as t2 on t1.project_id = t2.id where t2.uuid = %s
"""


class ProjectFile:
    def __init__(self, db):
        self.__db = db

    def get_by_id(self, id: int):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_FILE_GET_BY_ID, (id,),
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectFile(*i))
        return None if len(res) != 1 else res[0]

    def get_by_uuid(self, uuid):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_FILE_GET_BY_UUID, (uuid,),
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectFile(*i))
        return res

    def get_meta_by_project_id(self, project_id):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_FILE_GET_META_BY_PROJECT_ID, (project_id,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectFile(*i))
        return res

    def create(self, name, body, project_id, uuid):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_FILE_CREATE, (name, body, project_id, uuid,),
                )
        return self.get_by_uuid(uuid)

    def get_meta_by_project_uuid(self, project_uuid: str):
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_FILE_GET_META_BY_PROJECT_UUID, (project_uuid,),
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectFile(*i))
        return res
