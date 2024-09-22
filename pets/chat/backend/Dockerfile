# Используем базовый образ с поддержкой CUDA
FROM nvidia/cuda:11.8.0-cudnn8-devel-ubuntu22.04

# Установим системные зависимости
RUN apt-get update && \
    apt-get install -y \
    python3 python3-pip python3-dev \
    curl gnupg apt-transport-https ca-certificates \
    && apt-get clean

# Все необходимые библиотеки
RUN pip install flask && \
    pip install pika && \
    pip install pypdf && \
    pip install chromadb && \
    pip install langchain && \
    pip install Flask-Cors && \
    pip install python-docx && \
    pip install transformers && \
    pip install huggingface_hub && \
    pip install langchain_community && \
    pip install langchain_huggingface && \
    CMAKE_ARGS="-DLLAMA_CUBLAS=on" FORCE_CMAKE=1 pip install llama-cpp-python --force-reinstall --upgrade --no-cache-dir

# Копируем код приложения в контейнер
WORKDIR /app
COPY . /app

# Команда по умолчанию: запуск Flask сервера
CMD ["python3", "main.py", "server"]
