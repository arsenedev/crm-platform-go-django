from django import forms


class ContactCreateForm(forms.Form):

    customer_id = forms.ChoiceField(
        widget=forms.Select(
            attrs={
                "class": "form-select"
            }
        )
    )

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

    email = forms.EmailField(
        widget=forms.EmailInput(
            attrs={
                "class": "form-control"
            }
        )
    )

    phone = forms.CharField(
        required=False,
        widget=forms.TextInput(
            attrs={
                "class": "form-control"
            }
        )
    )