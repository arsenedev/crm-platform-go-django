from django.shortcuts import render
from django.shortcuts import redirect

from core.services.user_service import UserService

from .forms import (
    UserCreateForm,
    UserUpdateForm,
)


def users_list(request):

    token = request.session["token"]

    users = UserService.list(token)

    return render(
        request,
        "users/list.html",
        {
            "users": users,
        },
    )


def user_create(request):

    token = request.session["token"]

    form = UserCreateForm()

    if request.method == "POST":

        form = UserCreateForm(request.POST)

        if form.is_valid():
            print(form.cleaned_data)
            UserService.create(
                form.cleaned_data,
                token,
            )

            return redirect(
                "users_list"
            )

    return render(
        request,
        "users/create.html",
        {
            "form": form,
        },
    )


def user_edit(
    request,
    user_id,
):

    token = request.session["token"]

    user = UserService.get(
        user_id,
        token,
    )

    if request.method == "POST":

        form = UserUpdateForm(
            request.POST
        )

        if form.is_valid():

            UserService.update(
                user_id,
                form.cleaned_data,
                token,
            )

            return redirect(
                "users_list"
            )

    else:

        form = UserUpdateForm(
            initial={
                "firstname": user["firstname"],
                "lastname": user["lastname"],
                "role": user["role"],
                "active": user["active"],
            }
        )

    return render(
        request,
        "users/edit.html",
        {
            "form": form,
            "user": user,
        },
    )


def user_delete(
    request,
    user_id,
):

    token = request.session["token"]

    UserService.delete(
        user_id,
        token,
    )

    return redirect(
        "users_list"
    )