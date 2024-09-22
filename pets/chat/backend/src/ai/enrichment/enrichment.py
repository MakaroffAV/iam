from src.config.llm import LLM as ConfigLLM


class Enrichment:
    def __init__(self, config_llm: ConfigLLM) -> None:
        self.__llm = config_llm.llama

    def __build_promt(self, question: str, context: str) -> str:
        prompt = """
        Ответь лаконично на вопрос при наличии контекста. Если ответить невозможно, не пытайся выдумать ответ
        Вопрос:
            {}
        Контекст:
            {}
        """.format(
            question, context,
        )
        return prompt

    def do(self, question: str, context: str) -> str:
        response = self.__llm.create_chat_completion(
            messages=[
                {"role": "system", "content": self.__build_promt(question, context)},
            ]
        )
        return response["choices"][0]["message"]["content"]
