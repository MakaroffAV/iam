from transformers import AutoTokenizer
from langchain_community.llms import LlamaCpp


LLM_QUESTION_TEMPLATE = """
Задача:
    Ответь на вопрос при наличии контекста
Вопрос:
    {}
Контекст:
    {}
"""


class LLM:
    def __init__(self) -> None:
        self.__model_name = "openchat/openchat-3.6-8b-20240522"
        self.__model_path = "./model/openchat-3.6-8b-20240522-IQ4_XS.gguf"

        self.__tkn = AutoTokenizer.from_pretrained(self.__model_name)
        self.__llm = LlamaCpp(
            model_path=self.__model_path, temperatiure=1, max_tokens=200, top_p=1, top_k=1, repeat_penalty=1,
        )

    def query(self, question: str, context: str):
        print(self.__llm.invoke(self.__format_query(question, context)))

    def __format_query(self, question: str, context: str):
        msg = [{
            "role": "user", "content": LLM_QUESTION_TEMPLATE.format(question, context)
        }]
        return self.__tkn.apply_chat_template(msg, tokenize=False, add_generation_prompt=True)
