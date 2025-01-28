package memberships

import (
	"net/http"

	"github.com/NXRts/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SingUp(c *gin.Context) {
	var request memberships.SingUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.SingUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
