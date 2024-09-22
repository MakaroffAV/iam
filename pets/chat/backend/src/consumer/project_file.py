from modules.email import Email as ModuleEmail

from src.repo.user import User as RepoUser
from src.repo.project import Project as RepoProject
from src.repo.project_file import ProjectFile as RepoProjectFile
from src.repo.ch_database.chroma import Chroma as RepoChroma

from src.ai.loader.loader import Loader as LoaderAI
from src.config.embedding import Embedding as ConfigEmbedding


class ProjectFile:
    def __init__(
            self,
            embeddings: ConfigEmbedding,
            email_agent: ModuleEmail,
            rp_chroma_db: RepoChroma,
            rp_project: RepoProject,
            rp_project_file: RepoProjectFile,
            rp_user: RepoUser) -> None:

        self.__rp_project = rp_project
        self.__rp_project_file = rp_project_file
        self.__email_agent = email_agent
        self.__embeddings = embeddings
        self.__rp_chroma_db = rp_chroma_db
        self.__rp_user = rp_user

    def __get_text_data(self, project_id):
        data = {}
        project_files_meta \
            = self.__rp_project_file.get_meta_by_project_id(project_id)
        for i in project_files_meta:
            data[i.name] = LoaderAI(self.__rp_project_file.get_by_id(i.id)).do()
        return data

    def __set_project_ready(self, project_id):
        self.__rp_project.update_status_by_id(project_id, "Ready")

    def __set_project_failed(self, project_id):
        self.__rp_project.update_status_by_id(project_id, "Failed")

    def __send_project_is_ready(self, project_id):
        proj = self.__rp_project.get_by_id(project_id)[0]
        user = self.__rp_user.get_by_id(proj.user_id)[0]
        self.__email_agent.send_email(
            "Проект готов к работе: {}".format(proj.name), "Проект обновлен", user.email,
        )

    def __set_project_processing(self, project_id):
        self.__rp_project.update_status_by_id(project_id, "Processing")

    def __set_up_chroma_db(self, data: dict[str, list[str]], project_id):
        for i in self.__embeddings.models:
            self.__rp_chroma_db.create_collection(project_id, i["title"], data, i["model"])

    def do(self, project_id) -> None:
        try:
            self.__set_project_processing(project_id)

            project_text_data = self.__get_text_data(project_id)
            self.__set_up_chroma_db(project_text_data, project_id)

            self.__set_project_ready(project_id)
            self.__send_project_is_ready(project_id)
        except Exception as e:
            print(e)
            self.__set_project_failed(project_id)
