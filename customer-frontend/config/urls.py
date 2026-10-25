from django.contrib import admin
from django.urls import include
from django.urls import path

urlpatterns = [

    path(
        "admin/",
        admin.site.urls,
    ),

    path(
        "",
        include(
            "authentication.urls"
        ),
    ),

    path(
        "",
        include(
            "dashboard.urls"
        ),
    ),

    path(
        "users/",
        include(
            "users.urls"
        ),
    ),

    path(
        "customers/",
        include(
            "customers.urls"
        ),
    ),

    path(
        "contacts/",
        include(
            "contacts.urls"
        ),
    ),

    path(
        "audit/",
        include(
            "audit.urls"
        ),
    ),
]