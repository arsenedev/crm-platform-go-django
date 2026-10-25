from django.urls import path

from .views import (
    contacts_list,
    contact_create,
)

urlpatterns = [

    path(
        "",
        contacts_list,
        name="contacts_list",
    ),

    path(
        "create/",
        contact_create,
        name="contact_create",
    ),
]