import os


class Email:
    def __init__(self):
        self.host = os.getenv("MAIL_HOST")
        self.port = os.getenv("MAIL_PORT")
        self.user = os.getenv("MAIL_USER")
        self.pasw = os.getenv("MAIL_PASW")
