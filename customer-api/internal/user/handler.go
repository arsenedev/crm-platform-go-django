package user

import (
	"net/http"
	"strings"

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
func (h *Handler) CreateUser(
	c *gin.Context,
) {

	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.service.Create(
		c.Request.Context(),
		req,
	)

	if err != nil {

		if strings.Contains(
			err.Error(),
			"users_email_key",
		) {

			c.JSON(
				http.StatusConflict,
				gin.H{
					"error": "email already exists",
				},
			)

			return
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "user created",
		},
	)
}

func (h *Handler) ListUsers(
	c *gin.Context,
) {

	users, err := h.service.List(
		c.Request.Context(),
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
		users,
	)
}
func (h *Handler) GetUser(
	c *gin.Context,
) {

	user, err := h.service.GetByID(
		c.Request.Context(),
		c.Param("id"),
	)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "user not found",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		user,
	)
}
func (h *Handler) DeleteUser(
	c *gin.Context,
) {

	err := h.service.Delete(
		c.Request.Context(),
		c.Param("id"),
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
			"message": "user deleted",
		},
	)
}

func (h *Handler) UpdateUser(
	c *gin.Context,
) {

	id := c.Param("id")

	var req UpdateUserRequest

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
			"message": "user updated",
		},
	)
}
