from django.shortcuts import render
from django.shortcuts import redirect

from .forms import (
    ContactCreateForm,
    
)

from core.services.contact_service import (
    ContactService,
)

from core.services.customer_service import (
    CustomerService,
)


def contacts_list(request):

    token = request.session["token"]

    contacts = ContactService.list(
        token,
    )

    return render(
        request,
        "contacts/list.html",
        {
            "contacts": contacts,
        },
    )


def contact_create(request):

    token = request.session["token"]

    customers = CustomerService.list(token)

    customer_choices = [
        (
            customer["id"],
            customer["name"]
        )
        for customer in customers
    ]

    form = ContactCreateForm()

    form.fields["customer_id"].choices = customer_choices

    if request.method == "POST":

        form = ContactCreateForm(request.POST)

        form.fields["customer_id"].choices = customer_choices

        if form.is_valid():

            print(form.cleaned_data)

            ContactService.create(
                form.cleaned_data,
                token,
            )

            return redirect(
                "contacts_list"
            )

    return render(
        request,
        "contacts/create.html",
        {
            "form": form,
        },
    )
