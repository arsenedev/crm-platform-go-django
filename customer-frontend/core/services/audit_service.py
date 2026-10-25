from core.services.api_client import APIClient


class AuditService:

    @staticmethod
    def list(token):

        return APIClient.get(
            "/api/v1/audit",
            token,
        )