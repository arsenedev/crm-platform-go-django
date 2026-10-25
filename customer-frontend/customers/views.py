from django.shortcuts import render
from django.shortcuts import redirect

from .forms import (
    CustomerCreateForm,
    CustomerUpdateForm,
)

from core.services.customer_service import (
    CustomerService,
)


def customers_list(request):

    token = request.session["token"]

    customers = CustomerService.list(
        token
    )

    return render(
        request,
        "customers/list.html",
        {
            "customers": customers,
        },
    )


def customer_create(request):

    token = request.session["token"]

    form = CustomerCreateForm()

    if request.method == "POST":

        form = CustomerCreateForm(
            request.POST
        )

        if form.is_valid():

            print(form.cleaned_data)

            CustomerService.create(
                form.cleaned_data,
                token,
            )

            return redirect(
                "customers_list"
            )

    return render(
        request,
        "customers/create.html",
        {
            "form": form,
        },
    )


def customer_edit(
    request,
    customer_id,
):

    token = request.session["token"]

    customer = CustomerService.get(
        customer_id,
        token,
    )

    if request.method == "POST":

        form = CustomerUpdateForm(
            request.POST
        )

        if form.is_valid():

            print(form.cleaned_data)

            CustomerService.update(
                customer_id,
                form.cleaned_data,
                token,
            )

            return redirect(
                "customers_list"
            )

    else:

        form = CustomerUpdateForm(
            initial={
                "name": customer["name"],
                "email": customer["email"],
                "phone": customer["phone"],
                "status": customer["status"],
            }
        )

    return render(
        request,
        "customers/edit.html",
        {
            "form": form,
            "customer": customer,
        },
    )


def customer_delete(
    request,
    customer_id,
):

    token = request.session["token"]

    CustomerService.delete(
        customer_id,
        token,
    )

    return redirect(
        "customers_list"
    )