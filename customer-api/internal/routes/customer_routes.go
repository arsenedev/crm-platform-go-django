package routes

import (
	"company/customer-api/internal/auth"

	"github.com/gin-gonic/gin"
)

func registerCustomers(
	api *gin.RouterGroup,
	h Handlers,
) {

	customers := api.Group("/customers")

	customers.GET(
		"",
		h.Customer.ListCustomers)

	customers.GET(
		"/:id",
		h.Customer.GetCustomer,
	)

	customers.POST(
		"",
		h.Customer.CreateCustomer,
	)

	customers.PUT(
		"/:id",
		h.Customer.UpdateCustomer,
	)

	customers.DELETE(
		"/:id",
		auth.RequireRole("admin"),
		h.Customer.DeleteCustomer,
	)
}
