from django.shortcuts import render
from django.shortcuts import redirect

from core.services.user_service import UserService
from core.services.customer_service import CustomerService
from core.services.contact_service import ContactService
from core.services.audit_service import AuditService


def dashboard(request):

    token = request.session.get("token")

    if not token:
        return redirect("login")

    try:
        users = UserService.list(token) or []
    except Exception as e:
        print("Users error:", e)
        users = []

    try:
        customers = CustomerService.list(token) or []
    except Exception as e:
        print("Customers error:", e)
        customers = []

    try:
        contacts = ContactService.list(token) or []
    except Exception as e:
        print("Contacts error:", e)
        contacts = []

    try:
        audits = AuditService.list(token) or []
    except Exception as e:
        print("Audits error:", e)
        audits = []

    return render(
        request,
        "dashboard/index.html",
        {
            "users_count": len(users),
            "customers_count": len(customers),
            "contacts_count": len(contacts),
            "audit_count": len(audits),
        },
    )
