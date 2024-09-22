import json

from modules.email import Email as ModuleEmail

from src.config.llm import LLM as ConfigLLM
from src.config.email import Email as ConfigEmail
from src.config.db import Database as ConfigDatabase
from src.config.broker import Broker as ConfigBroker
from src.config.chroma import ChromaDB as ConfigChromaDB
from src.config.embedding import Embedding as ConfigEmbedding

from src.repo.user import User as RepoUser
from src.repo.project import Project as RepoProject
from src.repo.ch_database.chroma import Chroma as RepoChroma
from src.repo.project_file import ProjectFile as RepoProjectFile
from src.repo.project_query import ProjectQuery as RepoProjectQuery

from src.ai.enrichment.enrichment import Enrichment as AI_Enrichment

from src.consumer.project_file import ProjectFile as ConsumerProjectFile
from src.consumer.project_query import ProjectQuery as ConsumerProjectQuery


embeddings = ConfigEmbedding()

consumer_project_file = ConsumerProjectFile(
    embeddings, ModuleEmail(ConfigEmail()), RepoChroma(ConfigChromaDB()), RepoProject(ConfigDatabase()), RepoProjectFile(ConfigDatabase()), RepoUser(ConfigDatabase())
)

consumer_project_query = ConsumerProjectQuery(
    RepoChroma(ConfigChromaDB()), RepoProjectQuery(ConfigDatabase()), embeddings, RepoProject(ConfigDatabase()), RepoUser(ConfigDatabase()), ModuleEmail(ConfigEmail()), AI_Enrichment(ConfigLLM()),
)


def consumer():
    conn = ConfigBroker().connection
    chan = conn.channel()
    chan.basic_consume(
        auto_ack=True, queue="ai_app", on_message_callback=consumer_callback,
    )
    chan.start_consuming()


def consumer_callback(ch, meth, prop, body):
    body_json = json.loads(body)
    print(body_json)
    match body_json.get("task"):
        case None:
            print("unknown message type")
        case "create_project":
            try:
                consumer_project_file.do(
                    body_json["data"]["project"]["id"],
                )
            except Exception as e:
                print(e)
        case "new_project_query":
            try:
                consumer_project_query.do(
                    body_json["data"]["id"]
                )
            except Exception as e:
                print(e)
