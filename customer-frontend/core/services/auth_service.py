from core.services.api_client import (
    APIClient,
)


class AuthService:

    @staticmethod
    def login(
        email,
        password,
    ):
        return APIClient.login(
            email,
            password,
        )