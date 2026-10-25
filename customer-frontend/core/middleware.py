from django.shortcuts import redirect


class LoginRequiredMiddleware:

    def __init__(
        self,
        get_response,
    ):

        self.get_response = (
            get_response
        )

    def __call__(
        self,
        request,
    ):

        public_urls = [
            "/login/",
        ]

        if (
            request.path not in public_urls
            and not request.session.get(
                "token"
            )
        ):

            return redirect(
                "login"
            )

        return self.get_response(
            request
        )