import json
import requests

class RestRepository:

    def __init__(self, url:str = "http://localhost:8080"):
        self.url = url

    def registration(self, email:str, username:str, password:str):
        path = self.url + "/auth/sign-up"
        input_data = json.dumps({
            "email": email,
            "username": username,
            "password" : password
        })
        headers = {'Content-Type': "application/json; charset=utf-8"}

        response = requests.post(path, data=input_data, headers=headers)

        print(response.json())
        return response.status_code, response.json()

    def authorization(self, email:str, password:str):
        path = self.url + "/auth/sign-in"
        input_data = json.dumps({
            "email": email,
            "password" : password
        })
        headers = {'Content-Type': "application/json; charset=utf-8"}

        response = requests.post(path, data=input_data, headers=headers)

        return response.status_code, response.json()