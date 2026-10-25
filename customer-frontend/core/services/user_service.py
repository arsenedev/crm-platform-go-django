from core.services.api_client import APIClient


class UserService:

    @staticmethod
    def list(token):

        return APIClient.get(
            "/api/v1/users",
            token,
        )

    @staticmethod
    def get(user_id, token):

        return APIClient.get(
            f"/api/v1/users/{user_id}",
            token,
        )

    @staticmethod
    def create(data, token):

        return APIClient.post(
            "/api/v1/users",
            data,
            token,
        )

    @staticmethod
    def update(
        user_id,
        data,
        token,
    ):

        return APIClient.put(
            f"/api/v1/users/{user_id}",
            data,
            token,
        )

    @staticmethod
    def delete(
        user_id,
        token,
    ):

        return APIClient.delete(
            f"/api/v1/users/{user_id}",
            token,
        )