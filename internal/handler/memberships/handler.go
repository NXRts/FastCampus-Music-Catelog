package memberships

import (
	"github.com/NXRts/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

type service interface {
	SingUp(request memberships.SingUpRequest) error
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/memberships")
	route.POST("/sing_up", h.SingUp)
}
