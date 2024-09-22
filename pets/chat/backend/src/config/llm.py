import os

from llama_cpp import Llama
from huggingface_hub import login
from huggingface_hub import hf_hub_download


class LLM:
    def __init__(self):
        login(token=os.getenv("HF_TOKEN"),)
        login(token="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")

        model_basename = "openchat-3.6-8b-20240522-IQ4_XS.gguf"
        model_name_or_path = "bartowski/openchat-3.6-8b-20240522-GGUF"

        model_path = hf_hub_download(
            repo_id=model_name_or_path, filename=model_basename,
        )

        self.llama = Llama(
            model_path=model_path, n_gpu_layers=-1, n_ctx=4096, verbose=True,
        )
