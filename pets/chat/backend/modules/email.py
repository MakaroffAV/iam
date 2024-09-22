import smtplib

from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart


class Email:
    def __init__(self, config) -> None:
        self.__config = config

    def send_email(self, body: str, subject: str, recipient: str):
        with smtplib.SMTP(self.__config.host, self.__config.port) as smtp_server:
            smtp_server.starttls()
            smtp_server.login(
                self.__config.user,
                self.__config.pasw,
            )

            msg = MIMEMultipart()
            msg["To"] = recipient
            msg["From"] = self.__config.user
            msg["Subject"] = subject

            part = MIMEText(body, "html")
            msg.attach(part)
            text = msg.as_string()
            smtp_server.sendmail(self.__config.user, recipient, text)
