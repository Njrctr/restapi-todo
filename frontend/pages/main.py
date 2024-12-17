import flet as ft
from flet_route import Params, Basket
from repository.rest_repository import RestRepository
from utils import styles


class MainPage: #Главная страница


    db = RestRepository("http://localhost:8080")


    def view(self, page: ft.Page, params: Params, basket: Basket) -> ft.View:
        page.title = "Главная страница"
        page.fonts = {"zametka":"fonts/Zametka_Parletter.otf"}

        user_token = page.session.get('token')
        return ft.View(
            route='/main',
            controls=[
                ft.Text(f"hello from main page! Your token is: {user_token}")
            ],
            bgcolor=styles.defaultBgColor,
            padding=0
        )
