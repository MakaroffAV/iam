import os
import psycopg2


class Database:

    def __init__(self) -> None:
        self.__user = os.getenv("PDB_USER")
        self.__pass = os.getenv("PDB_PASS")
        self.__host = os.getenv("PDB_HOST")
        self.__port = os.getenv("PDB_PORT")
        self.__name = os.getenv("PDB_NAME")

    def __dsn(self) -> str:
        pass

    def __enter__(self):
        self.conn = psycopg2.connect(
            host=self.__host,
            port=self.__port,
            user=self.__user,
            password=self.__pass,
            database=self.__name,
        )
        return self.conn

    def __exit__(self, exc_type, exc_val, exc_tb):
        if not exc_type:
            self.conn.commit()
        else:
            self.conn.rollback()

        if self.conn: self.conn.close()
