from docxtpl import DocxTemplate


class QueryResponse:
    def __init__(self, query: str, project_name: str, use_enrichment: bool, query_created: str, enriched_query_response: str, results: dict[str, str], query_uuid) -> None:
        """
        :param results: 
            must be dict like this: 
                {"model": "model_name", "distance": 50, "result": "query_result", "source": "query_doc_source"}
        """

        self.__template = DocxTemplate("./template/template.docx")
        self.__query_uuid = query_uuid

        self.__context = {
            "query": query,
            "project_name": project_name,
            "use_enrichment": "Да" if use_enrichment else "Нет",
            "query_created": query_created,
            "enriched_query_response": enriched_query_response if enriched_query_response is not None else "Не используется",
            "results": results,
        }

    def render(self):
        self.__template.render(self.__context)
        self.__template.save(
            "./queries/{}.docx".format(self.__query_uuid),
        )
