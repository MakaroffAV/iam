import uuid
import random

from src.repo.user import User as RepoUser
from src.repo.code import Code as RepoCode

from modules.email import Email as ModuleEmail


class Login:
    
    def __init__(
            self,
            rp_user: RepoUser,
            rp_code: RepoCode,
            md_mail: ModuleEmail) -> None:

        self.__md_mail = md_mail
        self.__rp_user = rp_user
        self.__rp_code = rp_code

    def login(self, email: str):
        if not self.__rp_user.exists(email):
            user = self.__rp_user.create(email)
        else:
            user = self.__rp_user.get_by_email(email)

        otp = random.randint(1000, 9999)
        self.__md_mail.send_email(
            "Код подтверждения: {}".format(otp),
            "Код для входа в личный кабинет", user[0].email,
        )
        return self.__rp_code.create(user[0].id, otp, str(uuid.uuid4()))
