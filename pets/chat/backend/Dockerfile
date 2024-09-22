# Используем базовый образ с поддержкой CUDA
FROM nvidia/cuda:11.8.0-cudnn8-devel-ubuntu22.04

# Установим системные зависимости
RUN apt-get update && \
    apt-get install -y

# Все необходимые библиотеки
RUN pip3 install flask && \
    pip3 install pika && \
    pip3 install pypdf && \
    pip3 install chromadb && \
    pip3 install langchain && \
    pip3 install Flask-Cors && \
    pip3 install python-docx && \
    pip3 install transformers && \
    pip3 install huggingface_hub && \
    pip3 install langchain_community && \
    pip3 install langchain_huggingface && \
    CMAKE_ARGS="-DLLAMA_CUBLAS=on" FORCE_CMAKE=1 pip install llama-cpp-python --force-reinstall --upgrade --no-cache-dir

# Копируем код приложения в контейнер
WORKDIR /app
COPY . /app

# Команда по умолчанию: запуск Flask сервера
CMD ["python3", "main.py", "server"]
