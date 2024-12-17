import flet as ft
from router import Router

def main(page: ft.Page):
    start_page = "/signin"
    page.route = start_page
    Router(page)

if __name__== '__main__':
    ft.app(target=main)
