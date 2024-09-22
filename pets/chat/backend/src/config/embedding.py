from langchain_huggingface import HuggingFaceEmbeddings


class Embedding:
    def __init__(self):
        self.__device = "cuda"
        self.__cache_folder = "./embeddings"

        self.models = [
            {
                "title": "sber",
                "model": HuggingFaceEmbeddings(
                    model_name="inkoziev/sbert_pq", model_kwargs={"device": self.__device}, cache_folder=self.__cache_folder,
                ),
            },
            {
                "title": "LaBSE",
                "model": HuggingFaceEmbeddings(
                    model_name="cointegrated/LaBSE-en-ru", model_kwargs={"device": self.__device}, cache_folder=self.__cache_folder,
                ),
            },
            {
                "title": "rubert",
                "model": HuggingFaceEmbeddings(
                    model_name="cointegrated/rubert-tiny2", model_kwargs={"device": self.__device}, cache_folder=self.__cache_folder,
                ),
            },
        ]