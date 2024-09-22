from src.domain.project_query import ProjectQuery as DomainProjectQuery

SQL_PROJECT_QUERY_CREATE = """
insert into ai_app.project_query (query, uuid, project_id, enrichment) values (%s, %s, %s, %s)
"""

SQL_PROJECT_GET_BY_UUID = """
select id, uuid, project_id, query, enrichment, status, created from ai_app.project_query where uuid = %s
"""

SQL_PROJECT_QUERY_GET_BY_PROJECT_ID = """
select id, uuid, project_id, query, enrichment, status, created from ai_app.project_query where project_id = %s ORDER BY created DESC
"""

SQL_PROJECT_QUERY_GET_BY_ID = """
select id, uuid, project_id, query, enrichment, status, created from ai_app.project_query where id = %s
"""

SQL_PROJECT_QUERY_UPDATE = """
update ai_app.project_query set status = %s where id = %s
"""


class ProjectQuery:
    def __init__(self, db) -> None:
        self.__db = db

    def get_by_id(self, id) -> list[DomainProjectQuery]:
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_QUERY_GET_BY_ID, (id,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectQuery(*i))
        return res

    def update_status_by_id(self, id: int, status: str):
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_QUERY_UPDATE, (status, id,)
                )
        return self.get_by_id(id)

    def get_by_uuid(self, uuid) -> list[DomainProjectQuery]:
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_GET_BY_UUID, (uuid,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectQuery(*i))
        return res

    def get_by_project_id(self, project_id: int) -> list[DomainProjectQuery]:
        res = []
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_QUERY_GET_BY_PROJECT_ID, (project_id,)
                )
                for i in db_curs.fetchall():
                    res.append(DomainProjectQuery(*i))
        return res

    def create(self, project_id: int, query: str, enrichment: bool, uuid: str) -> list[DomainProjectQuery]:
        with self.__db as db_conn:
            with db_conn.cursor() as db_curs:
                db_curs.execute(
                    SQL_PROJECT_QUERY_CREATE, (query, uuid, project_id, enrichment,)
                )
        return self.get_by_uuid(uuid)
