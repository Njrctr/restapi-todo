import re
import time
from flet import Container, Text
from utils import styles

class Validator:
    
    def is_valid_email(self, error_field:Text, email:Container):
        pattern=r"^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$"
        if re.match(pattern, email.content.value) is None:
            self.validate_error("Некорректный Email", error_field, email)
            return False
        return True
    
    def is_valid_username(self, error_field:Text, username:Container):
        if any(char in ".,:;!_\"'@*-+()/#¤%&)" for char in username.content.value):
            self.validate_error("Поле username не может содержать спецсимволы", error_field, username)
            return False
        return True
    
    def is_valid_password(self, error_field:Text, password:Container):
        if len(password.content.value) < 5:
            self.validate_error("Минимальная длина пароля 5 символов", error_field, password)
            return False
        if not any(char.isdigit() for char in password.content.value):
            self.validate_error("Пароль должен содержать минимум одну букву и цифру", error_field, password)
            return False
        if not any(char in ".,:;!_\"'@*-+()/#¤%&)" for char in password.content.value):
            self.validate_error("Пароль должен содержать специальные символы", error_field, password)
            return False
        return True

    def is_valid_confpassword(self, error_field:Text, password:Container, confpassword:Container):
        if password.content.value != confpassword.content.value:
            self.validate_error("Пароли не совпадают", error_field, confpassword)
            return False
        return True
    

    @staticmethod
    def validate_error(msg:str, error_field:Text, container:Container=None):
        error_field.value = msg
        if container is not None:
            container.content.bgcolor=styles.inputBgErrorColor
            container.content.update()
        error_field.update()
        time.sleep(3)
        error_field.value = ""
        if container is not None:
            container.content.bgcolor=styles.secondaryBgColor
            container.content.update()
        error_field.update()
        return False