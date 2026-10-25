from django.shortcuts import render

from core.services.audit_service import (
    AuditService,
)


def audit_list(request):

    token = request.session["token"]

    logs = AuditService.list(
        token,
    )

    return render(
        request,
        "audit/list.html",
        {
            "logs": logs,
        },
    )