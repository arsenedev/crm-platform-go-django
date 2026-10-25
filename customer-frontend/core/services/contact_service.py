from core.services.api_client import (
    APIClient,
)


class ContactService:

    @staticmethod
    def list(token):
        return APIClient.get(
            "/api/v1/contacts",
            token,
        )

    @staticmethod
    def get(
        contact_id,
        token,
    ):
        return APIClient.get(
            f"/api/v1/contacts/{contact_id}",
            token,
        )

    @staticmethod
    def create(
        data,
        token,
    ):
        return APIClient.post(
            "/api/v1/contacts",
            data,
            token,
        )

    @staticmethod
    def update(
        contact_id,
        data,
        token,
    ):
        return APIClient.put(
            f"/api/v1/contacts/{contact_id}",
            data,
            token,
        )

    @staticmethod
    def delete(
        contact_id,
        token,
    ):
        return APIClient.delete(
            f"/api/v1/contacts/{contact_id}",
            token,
        )