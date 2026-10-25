from django.urls import path

from .views import (
    customers_list,
    customer_create,
    customer_edit,
    customer_delete,
)

urlpatterns = [

    path(
        "",
        customers_list,
        name="customers_list",
    ),

    path(
        "create/",
        customer_create,
        name="customer_create",
    ),

    path(
        "<str:customer_id>/edit/",
        customer_edit,
        name="customer_edit",
    ),

    path(
        "<str:customer_id>/delete/",
        customer_delete,
        name="customer_delete",
    ),
]