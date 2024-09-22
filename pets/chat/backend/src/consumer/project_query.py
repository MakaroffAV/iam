from src.config.embedding import Embedding as ConfigEmbedding
from src.repo.ch_database.chroma import Chroma as RepoCroma
from src.repo.project_query import ProjectQuery as RepoProjectQuery
from src.repo.project import Project as RepoProject
from src.domain.project import Project as DomainProject
from src.repo.user import User as RepoUser
from modules.email import Email as ModuleEmail
from src.domain.project_query import ProjectQuery as DomainProjectQuery

from modules.query_response import QueryResponse as ModuleQueryResponse
from src.ai.enrichment.enrichment import Enrichment as AI_Enrichment


class ProjectQuery:
    def __init__(
            self,
            repo_chroma: RepoCroma,
            repo_project_query: RepoProjectQuery,
            embeddings: ConfigEmbedding,
            repo_project: RepoProject,
            repo_user: RepoUser,
            email_agent: ModuleEmail,
            enrichment: AI_Enrichment) -> None:

        self.__repo_chroma = repo_chroma
        self.__repo_project_query = repo_project_query
        self.__embeddings = embeddings
        self.__repo_project = repo_project
        self.__repo_user = repo_user
        self.__email_agent = email_agent
        self.__enrichment = enrichment

    def __set_status_ready(self, project_query_id: int):
        self.__repo_project_query.update_status_by_id(project_query_id, "Ready")

    def __set_status_failed(self, project_query_id: int):
        self.__repo_project_query.update_status_by_id(project_query_id, "Failed")

    def __set_status_processing(self, project_query_id: int):
        self.__repo_project_query.update_status_by_id(project_query_id, "Processing")

    def __send_project_query_is_ready(self, project: DomainProject, query: DomainProjectQuery):
        user = self.__repo_user.get_by_id(project.user_id)
        self.__email_agent.send_email(
            "Ваш запрос обработан: {} \n Ссылка на результат: {}".format(
                query.query, "http://localhost:5001/api/project/query/result?uuid={}".format(query.uuid)
            ),
            "Запрос обработан",
            user[0].email,
        )

    def do(self, project_query_id: int) -> None:
        try:

            # Вытаскиваем метаданные запроса
            project_query = self.__repo_project_query.get_by_id(
                project_query_id,
            )

            # Обновляем статус,
            # что запрос пользователя в обработке
            self.__set_status_processing(project_query_id)

            # Выполняем поиск в векторной базе данных
            query_results = self.__repo_chroma.search(
                project_query[0].query, project_query[0].project_id, self.__embeddings.models,
            )

            # Вытаскиваем информацию о проекте
            project = self.__repo_project.get_by_id(project_query[0].project_id)

            # Проверяем, запросил ли пользователь обогащение ответа,
            # Если нет - просто формируем ответ пользователю, если - да, пропускаем все это
            # через языковую модель для обогащения ответа на вопрос через использование контекста
            enrichment = None
            if project_query[0].enrichment:
                enrichment = self.__enrichment.do(
                    project_query[0].query, " ".join([i["result"] for i in query_results])[:4096]
                )

            # Формируем ответ пользователю
            ModuleQueryResponse(
                project_query[0].query,
                project[0].name,
                project_query[0].enrichment,
                project_query[0].created.strftime("%Y-%m-%d %H:%M:%S"),
                enrichment,
                query_results,
                project_query[0].uuid,
            ).render()
            self.__send_project_query_is_ready(project[0], project_query[0])
            self.__set_status_ready(project_query_id)

        except Exception as e:
            print(e)
            self.__set_status_failed(project_query_id)
