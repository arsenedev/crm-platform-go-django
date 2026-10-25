from django.urls import path

from .views import (
    users_list,
    user_create,
    user_edit,
    user_delete,
)

urlpatterns = [

    path(
        "",
        users_list,
        name="users_list",
    ),

    path(
        "create/",
        user_create,
        name="user_create",
    ),

    path(
        "<str:user_id>/edit/",
        user_edit,
        name="user_edit",
    ),

    path(
        "<str:user_id>/delete/",
        user_delete,
        name="user_delete",
    ),
]