from django.shortcuts import render

from core.services.user_service import UserService
from core.services.customer_service import CustomerService
from core.services.contact_service import ContactService
from core.services.audit_service import AuditService


def dashboard(request):

    token = request.session["token"]

    users = UserService.list(token)

    customers = CustomerService.list(token)

    contacts = ContactService.list(token)

    audits = AuditService.list(token)

    return render(
    request,
    "dashboard/index.html",
    {
        "users_count": len(users or []),
        "customers_count": len(customers or []),
        "contacts_count": len(contacts or []),
        "audit_count": len(audits or []),
    },
)