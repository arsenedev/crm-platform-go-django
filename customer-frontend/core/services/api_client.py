import os
import requests

API_URL = os.getenv("CRM_API_URL")


class APIClient:

    @staticmethod
    def login(
        email,
        password,
    ):
        response = requests.post(
            f"{API_URL}/auth/login",
            json={
                "email": email,
                "password": password,
            },
            timeout=10,
        )

        response.raise_for_status()

        return response.json()

    @staticmethod
    def get(
        endpoint,
        token,
    ):
        response = requests.get(
            f"{API_URL}{endpoint}",
            headers={
                "Authorization": (
                    f"Bearer {token}"
                )
            },
            timeout=10,
        )

        response.raise_for_status()

        return response.json()

    @staticmethod
    def post(
        endpoint,
        data,
        token,
    ):
        response = requests.post(
            f"{API_URL}{endpoint}",
            json=data,
            headers={
                "Authorization": (
                    f"Bearer {token}"
                )
            },
            timeout=10,
        )

        response.raise_for_status()

        return response.json()

    @staticmethod
    def put(
        endpoint,
        data,
        token,
    ):
        response = requests.put(
            f"{API_URL}{endpoint}",
            json=data,
            headers={
                "Authorization": (
                    f"Bearer {token}"
                )
            },
            timeout=10,
        )

        response.raise_for_status()

        return response.json()

    @staticmethod
    def delete(
        endpoint,
        token,
    ):
        response = requests.delete(
            f"{API_URL}{endpoint}",
            headers={
                "Authorization": (
                    f"Bearer {token}"
                )
            },
            timeout=10,
        )

        response.raise_for_status()

        return response.json()