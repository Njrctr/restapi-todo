import flet as ft
from flet_route import Routing, path
from pages.auth import signin, signup
from pages import main
from utils.styles import defaultWindowWidth, defaultWindowHeight

class Router:
    def __init__(self, page: ft.Page):
        page.window.width=defaultWindowWidth
        page.window.height=defaultWindowHeight
        page.window.min_width=800
        page.window.min_height=500
        self.page = page
        self.app_routes= [
            path(url='/signin', clear=False, view=signin.SignInPage().view),
            path(url='/signup', clear=False, view=signup.SignUpPage().view),
            path(url='/main', clear=True, view=main.MainPage().view),
        ]
        
        Routing(
            page=self.page,
            app_routes=self.app_routes,
        )
        self.page.go(self.page.route)