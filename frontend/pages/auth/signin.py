import flet as ft
from flet_route import Params, Basket
from repository.rest_repository import RestRepository
from utils import styles, validate


class SignInPage: #Страница авторизации


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

    def view(self, page: ft.Page, params: Params, basket: Basket) -> ft.View:
        page.title = "Страница Авторизации"
        page.fonts = {"zametka":"fonts/Zametka_Parletter.otf"}

        signUpLink = ft.Container(
            ft.Text("Создать аккаунт", color=styles.defaultFontColor, size=15),
            alignment=ft.alignment.center,
            on_click=lambda e: page.go("/signup"),
        )

        def signin_button_click(e):
            self.email_input.content.value=self.email_input.content.value.strip()
            self.password_input.content.value=self.password_input.content.value.strip()

            if self.email_input.content.value and \
                self.password_input.content.value: # Проверка на пустые поля

                result = self.db.authorization(self.email_input.content.value, 
                                    self.password_input.content.value)
                print(result, result[1], type(result[1]))
                if result[0] != 200:
                    errors = {
                       "sql: no rows in result set":"Неверный email или пароль",
                    }
                    self.email_input.content.value=''
                    self.password_input.content.value=''
                    self.email_input.update()
                    self.password_input.update()
                    validate.Validator.validate_error(msg=errors[result[1]['message']], error_field=self.error_field)
                    return
                
                page.session.set('token', result[1]['token'])
                page.go("/main")
                




        return ft.View(
            route='/signin',
            controls=[
                ft.Row(
                    expand=True, 
                    controls=[
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
                                    self.password_input,
                                    ft.Container(
                                        ft.Text("Войти", color=styles.defaultFontColor, size=15),
                                        alignment=ft.alignment.center,
                                        height=40,
                                        bgcolor=styles.hoverBgColor,
                                        on_click=signin_button_click,
                                    ),
                                    signUpLink
                                ]
                            )
                        ),
                        ft.Container(
                            expand=3,
                            image_src='images/bg_login.webp',
                            image_fit=ft.ImageFit.COVER,
                            content=ft.Column(
                                alignment=ft.MainAxisAlignment.CENTER,
                                horizontal_alignment=ft.CrossAxisAlignment.CENTER,
                                controls=[
                                    ft.Icon(name=ft.icons.LOCK_PERSON_ROUNDED, color=styles.hoverBgColor,
                                            size=160),
                                    ft.Text("Авторизация", color=styles.hoverBgColor, size=20,
                                            weight=ft.FontWeight.BOLD)
                                ]
                            )
                        )
                        
                    ]
                )
            ],
            bgcolor=styles.defaultBgColor,
            padding=0
        )
