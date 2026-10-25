from django import forms


ROLE_CHOICES = [
    ("admin", "Admin"),
    ("manager", "Manager"),
    ("user", "User"),
]


class UserCreateForm(forms.Form):

    email = forms.EmailField(
        widget=forms.EmailInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    password = forms.CharField(
        widget=forms.PasswordInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    firstname = forms.CharField(
        widget=forms.TextInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    lastname = forms.CharField(
        widget=forms.TextInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    role = forms.ChoiceField(
        choices=[
            ("admin", "Admin"),
            ("manager", "Manager"),
            ("user", "User"),
        ],
        widget=forms.Select(
            attrs={
                "class": "form-select"
            }
        )
    )
class UserUpdateForm(forms.Form):

    firstname = forms.CharField(
        max_length=100,
        widget=forms.TextInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    lastname = forms.CharField(
        max_length=100,
        widget=forms.TextInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    role = forms.ChoiceField(
        choices=[
            ("admin", "Admin"),
            ("manager", "Manager"),
            ("user", "User"),
        ],
        widget=forms.Select(
            attrs={
                "class": "form-select"
            }
        )
    )

    active = forms.BooleanField(
        required=False,
        widget=forms.CheckboxInput(
            attrs={
                "class": "form-check-input"
            }
        )
    )