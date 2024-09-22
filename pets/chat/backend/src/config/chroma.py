import os

from chromadb import HttpClient


class ChromaDB:
    def __init__(self) -> None:
        self.client = HttpClient(
            host=os.getenv("CH_HOST"),
            port=os.getenv("CH_PORT"),
        )

    def search(self, query, model_name, project_uuid, model):
        query = model.embed_query(query)
        collection = self.__client.get_collection("{}_{}".format(project_uuid, model_name))
        results = collection.query(query_embeddings=[query], n_results=5)
        return [i for i in results["documents"][0]]

    def create(self, texts: list[str], model, uuid, model_name):
        collection = self.__client.create_collection(
            "{}_{}".format(uuid, model_name),
        )
        embeddings = model.embed_documents(texts)
        ids = [f"doc_{i}" for i in range(len(texts))]
        for i, v in enumerate(texts):
            collection.add(ids=[ids[i]], documents=[v], embeddings=[embeddings[i]])
