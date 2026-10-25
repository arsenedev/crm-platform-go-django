from core.services.api_client import (
    APIClient,
)


class CustomerService:

    @staticmethod
    def list(token):
        return APIClient.get(
            "/api/v1/customers",
            token,
        )

    @staticmethod
    def get(
        customer_id,
        token,
    ):
        return APIClient.get(
            f"/api/v1/customers/{customer_id}",
            token,
        )

    @staticmethod
    def create(
        data,
        token,
    ):
        return APIClient.post(
            "/api/v1/customers",
            data,
            token,
        )

    @staticmethod
    def update(
        customer_id,
        data,
        token,
    ):
        return APIClient.put(
            f"/api/v1/customers/{customer_id}",
            data,
            token,
        )

    @staticmethod
    def delete(
        customer_id,
        token,
    ):
        return APIClient.delete(
            f"/api/v1/customers/{customer_id}",
            token,
        )