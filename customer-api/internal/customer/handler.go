package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateCustomer(
	c *gin.Context,
) {

	var req CreateCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.service.Create(
		c.Request.Context(),
		req,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "customer created",
		},
	)
}

func (h *Handler) ListCustomers(
	c *gin.Context,
) {

	customers, err := h.service.List(
		c.Request.Context(),
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		customers,
	)
}

func (h *Handler) GetCustomer(
	c *gin.Context,
) {

	id := c.Param("id")

	customer, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "customer not found",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		customer,
	)
}

func (h *Handler) UpdateCustomer(
	c *gin.Context,
) {

	id := c.Param("id")

	var req UpdateCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.service.Update(
		c.Request.Context(),
		id,
		req,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "customer updated",
		},
	)
}

func (h *Handler) DeleteCustomer(
	c *gin.Context,
) {

	id := c.Param("id")

	err := h.service.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "customer deleted",
		},
	)
}
