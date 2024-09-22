import re
import pathlib

from io import BytesIO
from docx import Document
from pypdf import PdfReader

from langchain.text_splitter import RecursiveCharacterTextSplitter

from src.domain.project_file import ProjectFile as DomainProjectFile


class Loader:
    def __init__(
            self,
            project_file: DomainProjectFile) -> None:

        self.__chunk_size = 300
        self.__chunk_overlap = 20

        self.__project_file = project_file

    def do(self) -> list[str]:

        match pathlib.Path(self.__project_file.name).suffix:
            case ".pdf":
                return self.__handle_pdf()
            case ".txt":
                return self.__handle_txt()
            case ".docx":
                return self.__handle_doc()
            case _:
                return []

    def __splitter(self, text: str) -> list[str]:
        splitter = RecursiveCharacterTextSplitter(
            chunk_size=self.__chunk_size, chunk_overlap=self.__chunk_overlap,
        )
        return splitter.split_text(text)

    def __handle_txt(self):
        return self.__splitter(
            self.__project_file.body.tobytes().decode("utf-8"),
        )

    def __handle_doc(self):
        data = []
        reader = Document(BytesIO(self.__project_file.body))
        for i in reader.paragraphs:
            data.append(i.text)
        return self.__splitter(" ".join(data))

    def __handle_pdf(self):
        data = []
        reader = PdfReader(BytesIO(self.__project_file.body))
        for i in range(len(reader.pages)):
            data.append(
                re.sub(r'[\s\n]{2,}', '\n', reader.pages[i].extract_text().strip())
            )
        return self.__splitter(" ".join(data))