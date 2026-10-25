package routes

import (
	"github.com/gin-gonic/gin"
)

func registerContacts(
	api *gin.RouterGroup,
	h Handlers,
) {
	contacts := api.Group("/contacts")

	contacts.GET(
		"",
		h.Contact.ListContacts,
	)

	contacts.POST(
		"",
		h.Contact.CreateContact,
	)
}
