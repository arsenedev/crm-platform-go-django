from django.shortcuts import (
    render,
    redirect,
)

from .forms import LoginForm

from core.services.auth_service import (
    AuthService,
)


def login_view(request):

    if request.session.get(
        "token"
    ):
        return redirect(
            "dashboard"
        )

    form = LoginForm()

    if request.method == "POST":

        form = LoginForm(
            request.POST
        )

        if form.is_valid():

            try:

                result = AuthService.login(
                    form.cleaned_data["email"],
                    form.cleaned_data["password"],
                )

                request.session["token"] = (
                    result["token"]
                )

                request.session["email"] = (
                    form.cleaned_data["email"]
                )

                return redirect(
                    "dashboard"
                )

            except Exception:

                form.add_error(
                    None,
                    "Email ou mot de passe incorrect."
                )

    return render(
        request,
        "authentication/login.html",
        {
            "form": form,
        },
    )


def logout_view(request):

    request.session.flush()

    return redirect(
        "login"
    )