from langchain_huggingface import HuggingFaceEmbeddings
from src.config.chroma import ChromaDB as ConfigChromaDB


class Chroma:
    def __init__(self, ch_database: ConfigChromaDB) -> None:
        self.__ch_database = ch_database.client

    def search(
            self,
            query: str,
            project_id: str,
            embedding_models: list[dict[str, HuggingFaceEmbeddings]]) -> list[dict[str, str]]:

        results = []
        for i in embedding_models:
            embedded_query = i["model"].embed_query(query)
            collection_obj = self.__ch_database.get_collection(
                "{}_{}".format(project_id, i["title"])
            )
            query_results = collection_obj.query(
                query_embeddings=[embedded_query], n_results=3,
            )
            for j in range(len(query_results["documents"][0])):
                results.append(
                    {
                        "model": i["title"],
                        "source": query_results["ids"][0][j],
                        "distance": query_results["distances"][0][j],
                        "result": query_results["documents"][0][j],
                    }
                )

        return results

    def create_collection(
            self,
            project_id: str,
            embedding_model_name: str,
            texts: dict[str, list[str]],
            embedding_model: HuggingFaceEmbeddings) -> None:

        """
        :param texts:
            must be like this: <doc_name>: ["doc_chunk_1", "doc_chunk_2"]
        """

        collection = self.__ch_database.create_collection(
            "{}_{}".format(project_id, embedding_model_name)
        )

        # Собираем все тексты в один список
        # [text_1_chunk_1, ..., text_1_chunk_n, text_2_chunk_1, ..., text_2_chunk_n, ....]
        dat = []
        for i in texts:
            dat += texts[i]

        # Считаем эмбеддинги для текста
        embeddings = embedding_model.embed_documents(dat)

        # Собираем все докементы в один список
        # [doc_name_1_1, doc_name_1_2, ..., doc_name_2_1....]
        ids = []
        for i in texts:
            for j in range(len(texts[i])):
                ids.append("{}_{}".format(i, j))

        for i, v in enumerate(dat):
            collection.add(ids=[ids[i]], documents=[v], embeddings=[embeddings[i]])
