import time
import flet as ft
from flet_route import Params, Basket
from repository.rest_repository import RestRepository
from utils import styles
from utils.validate import Validator

class SignUpPage: # Страница регистрации

    valid = Validator()
    db = RestRepository("http://localhost:8080")

    error_field = ft.Text(" ", color="red")

    email_input=ft.Container(
        content=ft.TextField(
            label="Email",
            bgcolor=styles.secondaryBgColor,
            border=ft.InputBorder.NONE,
            filled=True,
            color=styles.secondaryFontColor
            ),
        border_radius=15
    )

    username_input=ft.Container(
        content=ft.TextField(
            label="Username",
            bgcolor=styles.secondaryBgColor,
            border=ft.InputBorder.NONE,
            filled=True,
            color=styles.secondaryFontColor
            ),
        border_radius=15
    )

    password_input=ft.Container(
        content=ft.TextField(
            label="Password", password=True, can_reveal_password=True,
            bgcolor=styles.secondaryBgColor,
            border=ft.InputBorder.NONE,
            filled=True,
            color=styles.secondaryFontColor
            ),
        border_radius=15
    )

    confPassword_input=ft.Container(
        content=ft.TextField(
            label="Confirm Password", password=True, can_reveal_password=True,
            bgcolor=styles.secondaryBgColor,
            border=ft.InputBorder.NONE,
            filled=True,
            color=styles.secondaryFontColor
            ),
        border_radius=15
    )

    def view(self, page: ft.Page, params: Params, basket: Basket):
        page.title= "Страница Регистрации"

        signInLink = ft.Container(
            ft.Text("Войти в аккаунт", color=styles.defaultFontColor, size=15),
            alignment=ft.alignment.center,
            on_click=lambda e: page.go("/signin"),
        )

        def signup_button_click(e):
            self.email_input.content.value=self.email_input.content.value.strip()
            self.username_input.content.value=self.username_input.content.value.strip()
            self.password_input.content.value=self.password_input.content.value.strip()
            self.confPassword_input.content.value=self.confPassword_input.content.value.strip()

            if self.email_input.content.value and self.username_input.content.value and \
                self.password_input.content.value and self.confPassword_input.content.value: # Проверка на пустые поля

                # todo Добавить проверку на Возможность добавить с данными email and username
                if self.valid.is_valid_email(self.error_field, self.email_input) and \
                    self.valid.is_valid_password(self.error_field, self.password_input) and \
                    self.valid.is_valid_confpassword(self.error_field, self.password_input, 
                                                     self.confPassword_input):
                    
                    result = self.db.registration(self.email_input.content.value, 
                                        self.username_input.content.value, 
                                        self.password_input.content.value)
                    print(result)
                    if result[0] == 200:
                        self.error_field.value = "Вы успешно зарегистрировались!"
                        self.error_field.color = "green"
                        self.error_field.update()
                        time.sleep(2)
                        self.error_field.value = ""
                        self.error_field.color = "red"
                        page.go("/signin")
                    else:
                        if result[1]['message'] == "pq: duplicate key value violates unique constraint \"users_username_key\"":
                            self.valid.validate_error("Пользователь с таким именем уже существует!", self.error_field)
                        elif result[1]['message'] == "pq: duplicate key value violates unique constraint \"users_email_key\"":
                            self.valid.validate_error("Пользователь с таким email уже существует!", self.error_field)
                    
            else:
                self.valid.validate_error("Заполните все поля", self.error_field)
            
        
        return ft.View(
            route='/signup',
            controls=[
                ft.Row(
                    expand=True, 
                    controls=[
                        ft.Container(
                            expand=3,
                            image_src='images/bg_login.webp',
                            image_fit=ft.ImageFit.COVER,
                            content=ft.Column(
                                alignment=ft.MainAxisAlignment.CENTER,
                                horizontal_alignment=ft.CrossAxisAlignment.CENTER,
                                controls=[
                                    ft.Icon(name=ft.icons.APP_REGISTRATION, color=styles.hoverBgColor,
                                            size=160),
                                    ft.Text("Регистрация", color=styles.hoverBgColor, size=20,
                                            weight=ft.FontWeight.BOLD)
                                ]
                            )
                        ),
                        ft.Container(
                            expand=2,
                            content=ft.Column(
                                alignment=ft.MainAxisAlignment.CENTER,
                                horizontal_alignment=ft.CrossAxisAlignment.CENTER,
                                controls=[
                                    ft.Text("Приветствую", color=styles.defaultFontColor, size=25, 
                                            weight=ft.FontWeight.NORMAL, font_family="zametka"
                                    ),
                                    self.error_field,
                                    self.email_input,
                                    self.username_input,
                                    self.password_input,
                                    self.confPassword_input,
                                    ft.Container(
                                        ft.Text("Зарегистрироваться", color=styles.defaultFontColor, size=15),
                                        alignment=ft.alignment.center,
                                        height=40,
                                        bgcolor=styles.hoverBgColor,
                                        on_click=signup_button_click,
                                    ),
                                    signInLink
                                ]
                            )
                        )
                        
                    ]
                )
            ],
            bgcolor=styles.defaultBgColor,
            padding=0
        )
    
    